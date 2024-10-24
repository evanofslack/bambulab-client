package monitor

import (
	"context"
	"time"

	mqtt "github.com/evanofslack/bambulab-client/mqtt"
)

// Monitor offers an abstracted view of a printer's state.
type Monitor struct {
	LastUpdate     time.Time
	Update         chan struct{}
	PrintStarted   chan struct{}
	PrintCancelled chan struct{}
	PrintFinished  chan struct{}
	messageHistory *messageHistory
	stateHistory   *stateHistory
	ctx            context.Context
	cancel         context.CancelFunc
}

// New creates a new monitor
func New() *Monitor {
	ctx, cancel := context.WithCancel(context.Background())
	m := &Monitor{
		Update:         make(chan struct{}),
		messageHistory: newMessageHistory(),
		stateHistory:   newStateHistory(),
		PrintStarted:   make(chan struct{}),
		PrintCancelled: make(chan struct{}),
		PrintFinished:  make(chan struct{}),
		ctx:            ctx,
		cancel:         cancel,
	}
	return m
}

// Start starts the monitor updating from the incoming messages
func (m *Monitor) Start(msgs <-chan mqtt.Message) {
	for {
		select {
		case <-m.ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				return
			}
			newMsg, changed := mergeMessage(m.messageHistory.current, &msg)
			if changed {
				m.handleChange(newMsg)
			}
		}
	}
}

// CurrentMessage is the current mqtt message in its raw form,
// same form as it comes from mqtt messages.
func (m *Monitor) CurrentMessage() mqtt.Message {
	return *m.messageHistory.current
}

// PreviousMessage is the previous mqtt message in its raw form,
// same form as it comes from mqtt messages.
func (m *Monitor) PreviousMessage() mqtt.Message {
	return *m.messageHistory.previous
}

// CurrentState is the current interpreted state
func (m *Monitor) CurrentState() State {
	return m.stateHistory.current
}

// PreviousState is the previous interpreted state
func (m *Monitor) PreviousState() State {
	return m.stateHistory.previous
}

func (m *Monitor) Stop() {
	m.cancel()
}

func (m *Monitor) handleChange(newMsg *mqtt.Message) {
	m.LastUpdate = time.Now()
	// Update history
	m.messageHistory.previous = m.messageHistory.current
	m.messageHistory.current = newMsg

	// Translate to state
	newState := stateFromMessage(newMsg)
	m.stateHistory.previous = m.stateHistory.current
	m.stateHistory.current = newState

	select {
	case <-m.ctx.Done():
		return
	case m.Update <- struct{}{}:
	default:
	}

	if isPrintStarted(m.stateHistory.current, m.stateHistory.previous) {
		select {
		case <-m.ctx.Done():
			return
		case m.PrintStarted <- struct{}{}:
		default:
		}
	}

	if isPrintCancelled(m.stateHistory.current) {
		select {
		case <-m.ctx.Done():
			return
		case m.PrintCancelled <- struct{}{}:
		default:
		}
	}

	if isPrintFinished(m.stateHistory.current, m.stateHistory.previous) {
		select {
		case <-m.ctx.Done():
			return
		case m.PrintFinished <- struct{}{}:
		default:
		}
	}
}

func isPrintStarted(curr, prev State) bool {
	if curr.Gcode.State.IsNone() || prev.Gcode.State.IsNone() {
		return false
	}
	cstate, pstate := curr.Gcode.State.Unwrap(), prev.Gcode.State.Unwrap()
	wasIdle := pstate == "IDLE" || pstate == "FAILED" || pstate == "FINISH"
	isIdle := cstate == "IDLE" || pstate == "FAILED" || pstate == "FINISH"
	started := wasIdle && !isIdle
	return started
}

func isPrintFinished(curr, prev State) bool {
	if curr.Gcode.State.IsNone() {
		return false
	}
	cstate := curr.Gcode.State.Unwrap()
	if cstate == "FINISH" && prev.Gcode.State.IsNone() {
		return true
	}
	if prev.Gcode.State.IsNone() {
		return false
	}
	pstate := prev.Gcode.State.Unwrap()
	finished := cstate == "FINISH" && pstate != "FINISH"
	return finished
}

func isPrintCancelled(curr State) bool {
	if curr.CurrentPrint.PrintError.IsNone() {
		return false
	}
	cerr := curr.CurrentPrint.PrintError.Unwrap()
	cancelled := cerr == 50348044
	return cancelled
}

type messageHistory struct {
	current  *mqtt.Message
	previous *mqtt.Message
}

func newMessageHistory() *messageHistory {
	m := &messageHistory{
		current:  &mqtt.Message{},
		previous: &mqtt.Message{},
	}
	return m
}

type stateHistory struct {
	current  State
	previous State
}

func newStateHistory() *stateHistory {
	s := &stateHistory{
		current:  State{},
		previous: State{},
	}
	return s
}
