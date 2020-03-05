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
			for _, c := range t.c {
				go slog.Recover(ctx, func(ctx context.Context) error {
					func(c chan<- KeyWithContexter, v KeyWithContexter) {
						const op = op + "#func"
						select {
						case <-ctx.Done():
							if err := ctx.Err(); err != nil {
								slog.Logger().Err(err).Str("op", op).Msg(err.Error())
							}
							return
						case c <- v:
						}
					}(c, v)
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
	mu sync.RWMutex
	m  map[interface{}]*tKeyWithContexter
}

func newMapKeyWithContexter() *mapKeyWithContexter {
	return &mapKeyWithContexter{
		m: make(map[interface{}]*tKeyWithContexter),
	}
}

func (m *mapKeyWithContexter) get(topic interface{}) TopicKeyWithContexter {
	const op = op + ".mapKeyWithContexter.get"

	if topic == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'topic' is nil"})
	}

	m.mu.RLock()
	v, ok := m.m[topic]
	m.mu.RUnlock()
	if ok {
		return v
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok = m.m[topic]
	if ok {
		return v
	}
	v = newTKeyWithContexter()
	m.m[topic] = v
	return v
}
