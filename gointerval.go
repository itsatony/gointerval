package gointerval

import "time"

// creates a new GoInteval struct that allows running a function on a regular interval.
// the call function can trigger a stop of the timer by returning anything but true
func NewInterval(call func() bool, duration time.Duration, runImmediately bool) *GoInterval {
	var iv GoInterval = GoInterval{
		tickDuration:     duration,
		call:             call,
		cancelChan:       make(chan bool),
		callHasCancelled: false,
	}
	iv.Start(duration, runImmediately)
	return &iv
}

type GoInterval struct {
	active           bool
	callHasCancelled bool
	tickDuration     time.Duration
	call             func() bool
	ticker           *time.Ticker
	cancelChan       chan (bool)
}

// Start (re-)starts the interval timer
func (iv *GoInterval) Start(duration time.Duration, runImmediately bool) *GoInterval {
	if iv.ticker != nil {
		iv.ticker.Stop()
	}
	iv.tickDuration = duration
	iv.ticker = time.NewTicker(iv.tickDuration)
	iv.active = true
	go func() {
		for {
			select {
			case <-iv.cancelChan:
				return
			case <-iv.ticker.C:
				if iv.call != nil {
					if !iv.call() {
						if iv.active {
							iv.callHasCancelled = true
							iv.Stop()
						}
						return
					}
				}
			}
		}
	}()
	if runImmediately {
		if !iv.call() {
			iv.Stop()
		}
	}
	return iv
}

// Stop stops the interval timer and resets it to the original duration
func (iv *GoInterval) Stop() *GoInterval {
	if !iv.active {
		return iv
	}
	iv.ticker.Stop()
	if !iv.callHasCancelled {
		iv.cancelChan <- true
	}
	iv.ticker.Reset(iv.tickDuration)
	iv.active = false
	return iv
}

// State returns the current state of the interval timer (true = running, false = stopped)
func (iv *GoInterval) State() bool {
	return iv.active
}
