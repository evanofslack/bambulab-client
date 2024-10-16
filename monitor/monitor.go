package monitor

import (
	"context"
	"time"

	mqtt "github.com/evanofslack/bambulab-client/mqtt"
)

type Monitor struct {
	State      *mqtt.Message
	Update     chan struct{}
	LastUpdate time.Time
	ctx        context.Context
	cancel     context.CancelFunc
}

func New() *Monitor {
	ctx, cancel := context.WithCancel(context.Background())
	m := &Monitor{
		State:  &mqtt.Message{},
		Update: make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
	}
	return m
}

func (m *Monitor) Start(msgs <-chan mqtt.Message) {
	for {
		select {
		case <-m.ctx.Done():
			return
		case msg, ok := <-msgs:
		    if !ok {
		        return
		    }
		    // fmt.Println("got new msg, check if state change")
            var changed bool
            m.State, changed = mergeState(m.State, &msg)
            if changed {
                // fmt.Println("state changed...")
				m.LastUpdate = time.Now()
				select {
				case <-m.ctx.Done():
					return
				// Send nonblocking update
				case m.Update <- struct{}{}:
				default:
				}
			}
		}
	}
}

func (m *Monitor) Stop() {
	m.cancel()
}
