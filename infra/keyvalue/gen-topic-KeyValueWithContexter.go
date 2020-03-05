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
	topicKeyValueWithContexter *mapKeyValueWithContexter
)

func init() {
	topicKeyValueWithContexter = newMapKeyValueWithContexter()
}

// GetTopicKeyValueWithContexter returns a TopicKeyValueWithContexter of the given topic.
func GetTopicKeyValueWithContexter(topic interface{}) TopicKeyValueWithContexter {
	return topicKeyValueWithContexter.get(topic)
}

// TopicKeyValueWithContexter is a topic.
type TopicKeyValueWithContexter interface {
	// Publish returns a '<-chan KeyValueWithContexter' that joins to the given topic.
	Publish(ctx context.Context, c <-chan KeyValueWithContexter)
	// Publisher returns a 'chan<- KeyValueWithContexter' that joins to the given topic.
	Publisher(ctx context.Context) chan<- KeyValueWithContexter
	// Subscribe returns a 'chan<- KeyValueWithContexter' that joins to the given topic.
	Subscribe(c chan<- KeyValueWithContexter)
}

type tKeyValueWithContexter struct {
	mu sync.RWMutex
	c  []chan<- KeyValueWithContexter
}

func newTKeyValueWithContexter() *tKeyValueWithContexter {
	return &tKeyValueWithContexter{
		c: make([]chan<- KeyValueWithContexter, 0),
	}
}

func (t *tKeyValueWithContexter) Publish(ctx context.Context, c <-chan KeyValueWithContexter) {
	const op = op + ".tKeyValueWithContexter.Publish"

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
					func(c chan<- KeyValueWithContexter, v KeyValueWithContexter) {
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

func (t *tKeyValueWithContexter) Publisher(ctx context.Context) chan<- KeyValueWithContexter {
	const op = op + ".tKeyValueWithContexter.Publisher"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}

	c := make(chan KeyValueWithContexter)
	t.Publish(ctx, c)
	return c
}

func (t *tKeyValueWithContexter) Subscribe(c chan<- KeyValueWithContexter) {
	const op = op + ".tKeyValueWithContexter.Subscribe"

	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	t.mu.Lock()
	t.c = append(t.c, c)
	t.mu.Unlock()
}

type mapKeyValueWithContexter struct {
	mu sync.RWMutex
	m  map[interface{}]*tKeyValueWithContexter
}

func newMapKeyValueWithContexter() *mapKeyValueWithContexter {
	return &mapKeyValueWithContexter{
		m: make(map[interface{}]*tKeyValueWithContexter),
	}
}

func (m *mapKeyValueWithContexter) get(topic interface{}) TopicKeyValueWithContexter {
	const op = op + ".mapKeyValueWithContexter.get"

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
	v = newTKeyValueWithContexter()
	m.m[topic] = v
	return v
}
