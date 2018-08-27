package bus

import (
	"fmt"
	"sync"

	"github.com/michilu/message-bus"

	"github.com/michilu/boilerplate/v/errs"
	"github.com/michilu/boilerplate/v/log"
)

var (
	bus messagebus.MessageBus
	wg  sync.WaitGroup

	register = make(map[string][]interface{})

	// Publish publishes arguments to the given topic subscribers.
	Publish func(topic string, args ...interface{})
)

func init() {
	bus = messagebus.New()
	Publish = bus.Publish
}

func Register(topic fmt.Stringer, fn interface{}) {
	const op = "bus.Register"
	k := topic.String()
	err := Subscribe(k, fn)
	if err != nil {
		log.Logger().Fatal().
			Str("op", op).
			Err(&errs.Error{Op: op, Err: err}).
			Msg("error")
	}
	register[k] = append(register[k], fn)
}

// Subscribe subscribes to the given topic.
func Subscribe(topic string, fn interface{}) error {
	const op = "bus.Subscribe"

	log.Logger().Debug().
		Str("op", op).
		Str("topic", topic).
		Msg("start")

	err := bus.Subscribe(topic, fn)
	if err != nil {
		return &errs.Error{Op: op, Err: err}
	}
	wg.Add(1)

	log.Logger().Debug().
		Str("op", op).
		Str("topic", topic).
		Msg("end")

	return nil
}

// Unsubscribe unsubsribes from the given topic.
func Unsubscribe(topic string, fn interface{}) error {
	const op = "bus.Unsubscribe"
	defer wg.Done()

	log.Logger().Debug().
		Str("op", op).
		Str("topic", topic).
		Msg("start")

	err := bus.Unsubscribe(topic, fn)
	if err != nil {
		return &errs.Error{Op: op, Err: err}
	}

	log.Logger().Debug().
		Str("op", op).
		Str("topic", topic).
		Msg("end")

	return nil
}

// Wait waits until unsubscribe all subscribers.
func Wait() {
	wg.Wait()
}
