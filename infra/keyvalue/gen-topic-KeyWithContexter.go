// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package keyvalue

import (
	"context"
	"sync"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"google.golang.org/grpc/codes"
)

var (
	topicKeyWithContexter *mapKeyWithContexter
)

func init() {
	topicKeyWithContexter = newMapKeyWithContexter()
}

// GetTopicKeyWithContexter returns a TopicKeyWithContexter of the given topic.
func GetTopicKeyWithContexter(topic interface{}) TopicKeyWithContexter {
	return topicKeyWithContexter.get(topic)
}

// TopicKeyWithContexter is a topic.
type TopicKeyWithContexter interface {
	// Publish returns a '<-chan KeyWithContexter' that joins to the given topic.
	Publish(ctx context.Context, c <-chan KeyWithContexter)
	// Publisher returns a 'chan<- KeyWithContexter' that joins to the given topic.
	Publisher(ctx context.Context) chan<- KeyWithContexter
	// Subscribe returns a 'chan<- KeyWithContexter' that joins to the given topic.
	Subscribe(c chan<- KeyWithContexter)
}

type tKeyWithContexter struct {
	mu sync.RWMutex
	c  []chan<- KeyWithContexter
}

func newTKeyWithContexter() *tKeyWithContexter {
	return &tKeyWithContexter{
		c: make([]chan<- KeyWithContexter, 0),
	}
}

func (t *tKeyWithContexter) Publish(ctx context.Context, c <-chan KeyWithContexter) {
	const op = op + ".tKeyWithContexter.Publish"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	go slog.Recover(ctx, func(ctx context.Context) error {
		const op = op + "#go"
	loop:
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return &errs.Error{Op: op, Err: err}
			}
			return nil
		default:
		}
		for v := range c {
			v0 := v
			for _, v := range t.c {
				v1 := v
				go slog.Recover(ctx, func(ctx context.Context) error {
					func(c chan<- KeyWithContexter, v KeyWithContexter) {
						const op = op + "#func"
						select {
						case <-ctx.Done():
							err := &errs.Error{Op: op, Code: codes.Canceled, Err: ctx.Err()}
							slog.Logger().Debug().Err(err).Str("op", op).Msg(err.Error())
							return
						case c <- v:
						}
					}(v1, v0)
					return nil
				})
			}
			goto loop
		}
		return nil
	})

}

func (t *tKeyWithContexter) Publisher(ctx context.Context) chan<- KeyWithContexter {
	const op = op + ".tKeyWithContexter.Publisher"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}

	c := make(chan KeyWithContexter)
	t.Publish(ctx, c)
	return c
}

func (t *tKeyWithContexter) Subscribe(c chan<- KeyWithContexter) {
	const op = op + ".tKeyWithContexter.Subscribe"

	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	t.mu.Lock()
	t.c = append(t.c, c)
	t.mu.Unlock()
}

type mapKeyWithContexter struct {
	m sync.Map
}

func newMapKeyWithContexter() *mapKeyWithContexter {
	return &mapKeyWithContexter{}
}

func (m *mapKeyWithContexter) get(topic interface{}) TopicKeyWithContexter {
	const op = op + ".mapKeyWithContexter.get"

	if topic == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'topic' is nil"})
	}

	v0, ok := m.m.Load(topic)
	if ok {
		if v1, ok := v0.(TopicKeyWithContexter); ok {
			return v1
		}
	}

	v2 := newTKeyWithContexter()
	m.m.Store(topic, v2)
	return v2
}
