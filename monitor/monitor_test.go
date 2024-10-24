package monitor

import (
	"context"
	"sync"
	"testing"
	"time"

	mqtt "github.com/evanofslack/bambulab-client/mqtt"
	"github.com/stretchr/testify/assert"
)

var (
	// Create mock messages to simulate print start, finish, and cancellation.
	stateIdle     = "IDLE"
	stateRunning  = "RUNNING"
	stateFinished = "FINISH"
	stateFailed   = "FAILED"
	msgIdle       = newStateMsg(stateIdle)
	msgRunning    = newStateMsg(stateRunning)
	msgFinished   = newStateMsg(stateFinished)
	msgFailed     = newStateMsg(stateFailed)
	msgCancelled  = newCancelMsg()
)

// TestNewMonitor tests the creation of a new Monitor.
func TestNewMonitor(t *testing.T) {
	monitor := New()

	assert.NotNil(t, monitor)
	assert.NotNil(t, monitor.Update)
	assert.NotNil(t, monitor.messageHistory)
	assert.NotNil(t, monitor.stateHistory)
	assert.NotNil(t, monitor.ctx)
	assert.NotNil(t, monitor.cancel)
}

// TestMonitor_Start tests that the monitor updates its state when receiving new messages.
func TestMonitor_Start(t *testing.T) {
	monitor := New()
	defer monitor.Stop()

	msgs := make(chan mqtt.Message, 1)

	// Create first to send.
	state := "IDLE"
	msg := newStateMsg(state)
	go monitor.Start(msgs)

	msgs <- msg
	time.Sleep(100 * time.Millisecond)
	// Assert that the monitor has updated its message history.
	assert.Equal(t, state, *monitor.messageHistory.current.Print.GcodeState)
}

// TestMonitor_HandleChange tests the Monitor's handleChange function.
func TestMonitor_HandleChange(t *testing.T) {
	monitor := New()
	defer monitor.Stop()

	state1, state2 := "IDLE", "RUNNING"
	msg1 := newStateMsg(state1)
	msg2 := newStateMsg(state2)

	// Simulate change in message.
	monitor.handleChange(&msg1)
	assert.Equal(t, state1, *monitor.messageHistory.current.Print.GcodeState)

	// Simulate another change in message.
	monitor.handleChange(&msg2)
	assert.Equal(t, state2, *monitor.messageHistory.current.Print.GcodeState)
	assert.Equal(t, state1, *monitor.messageHistory.previous.Print.GcodeState)
}

// TestMonitor_Stop tests that the monitor stops correctly.
func TestMonitor_Stop(t *testing.T) {
	monitor := New()
	monitor.Stop()

	select {
	case <-monitor.ctx.Done():
		// Success, the context was cancelled.
	default:
		t.Error("Expected the monitor context to be cancelled")
	}
}

func TestMonitor_PrintStart(t *testing.T) {
	ctx := context.Background()
	monitor := New()
	defer monitor.Stop()

	readyChan := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	// Must start validation thread listening on chan
	go func() {
		defer wg.Done()
		// communicate back validation thread is ready
		if err := signalReady(ctx, readyChan); err != nil {
			t.Error(err)
		}

		// listen for print started signal
		select {
		case <-ctx.Done():
			t.Error("expected print started event to be triggered")
			return
		case <-monitor.PrintStarted:
		}
	}()
	// wait for validation thread to be ready
	if err := waitReady(ctx, readyChan); err != nil {
		t.Error(err)
	}
	// Simulate print start.
	monitor.handleChange(&msgIdle)    // Initial state.
	monitor.handleChange(&msgRunning) // Transition to running.
	wg.Wait()
}

func TestMonitor_PrintFinish(t *testing.T) {
	ctx := context.Background()
	monitor := New()
	defer monitor.Stop()

	readyChan := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// communicate back validation thread is ready
		if err := signalReady(ctx, readyChan); err != nil {
			t.Error(err)
		}

		select {
		case <-ctx.Done():
			t.Error("expected print finished event to be triggered")
			return
		case <-monitor.PrintFinished:
		}
	}()
	// wait for validation thread to be ready
	if err := waitReady(ctx, readyChan); err != nil {
		t.Error(err)
	}
	// Simulate print finish.
	monitor.handleChange(&msgFinished)
	wg.Wait()
}

func TestMonitor_PrintCancel(t *testing.T) {
	ctx := context.Background()
	monitor := New()
	defer monitor.Stop()

	readyChan := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// communicate back validation thread is ready
		if err := signalReady(ctx, readyChan); err != nil {
			t.Error(err)
		}

		select {
		case <-ctx.Done():
			t.Error("expected print cancelled event to be triggered")
			return
		case <-monitor.PrintCancelled:
		}
	}()
	// wait for validation thread to be ready
	if err := waitReady(ctx, readyChan); err != nil {
		t.Error(err)
	}
	// Simulate print cancellation.
	monitor.handleChange(&msgCancelled)

	wg.Wait()
}

func TestMonitor_PrintFail(t *testing.T) {
	ctx := context.Background()
	monitor := New()
	defer monitor.Stop()

	readyChan := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// communicate back validation thread is ready
		if err := signalReady(ctx, readyChan); err != nil {
			t.Error(err)
		}

		select {
		case <-ctx.Done():
			t.Error("expected print cancelled event to be triggered")
			return
		case <-monitor.PrintFailed:
		}
	}()
	// wait for validation thread to be ready
	if err := waitReady(ctx, readyChan); err != nil {
		t.Error(err)
	}
	// Simulate print cancellation.
	monitor.handleChange(&msgFailed)

	wg.Wait()
}

func signalReady(ctx context.Context, r chan struct{}) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case r <- struct{}{}:
		return nil
	}
}

func waitReady(ctx context.Context, r chan struct{}) error {
	// wait for validation thread to be ready
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-r:
		return nil
	}
}

func newStateMsg(state string) mqtt.Message {
	msg := mqtt.Message{
		Print: &mqtt.Print{
			GcodeState: &state,
		},
	}
	return msg
}

func newCancelMsg() mqtt.Message {
	e := 50348044
	msg := mqtt.Message{
		Print: &mqtt.Print{
			PrintError: &e,
		},
	}
	return msg
}
