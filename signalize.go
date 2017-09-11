package signalize

import (
	"os"
	"errors"
	"os/signal"
)

type catcher func(signal os.Signal)

var waits []os.Signal
var hooks map[os.Signal]catcher
var exits map[os.Signal]bool

var hook_chan chan os.Signal
var exit_chan chan bool
var launched bool

func Catch(fn catcher, signals ...os.Signal)  {
	if hooks==nil {
		hooks = map[os.Signal]catcher{}
	}
	for _, sig := range signals{
		listen(sig)
		hooks[sig]=fn
	}
}

func Stops(signals ...os.Signal)  {
	if exits==nil {
		exits = map[os.Signal]bool{}
	}
	for _, sig := range signals{
		listen(sig)
		exits[sig]=true
	}
}

func Hook(block bool) error {
	if launched {
		return errors.New("already launched")
	}
	launched = true
	hook_chan = make(chan os.Signal, 1)
	exit_chan = make(chan bool, 1)
	signal.Notify(hook_chan, waits...)

	go func() {
		for {
			sig := <-hook_chan
			if fn, ok := hooks[sig]; ok {
				fn(sig)
			}
			if _, ok := exits[sig]; ok {
				exit_chan<-true
				break
			}
		}
	}()

	if block {
		<-exit_chan
	}
	return nil
}

func listen(signal os.Signal)  {
	if waits==nil {
		waits = make([]os.Signal, 0)
	}
	waits = append(waits, signal)
}
