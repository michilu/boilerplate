// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package domain

import (
	"context"
	"sync"

	"github.com/michilu/boilerplate/v/errs"
	"google.golang.org/grpc/codes"
)

var (
	topicDurationer *mapDurationer
)

func init() {
	topicDurationer = newMapDurationer()
}

// GetTopicDurationer returns a TopicDurationer of the given topic.
func GetTopicDurationer(topic interface{}) TopicDurationer {
	return topicDurationer.get(topic)
}

// TopicDurationer is a topic.
type TopicDurationer interface {
	// Publish returns a '<-chan Durationer' that joins to the given topic.
	Publish(ctx context.Context, c <-chan Durationer)
	// Publisher returns a 'chan<- Durationer' that joins to the given topic.
	Publisher(ctx context.Context) chan<- Durationer
	// Subscribe returns a 'chan<- Durationer' that joins to the given topic.
	Subscribe(c chan<- Durationer)
}

type tDurationer struct {
	mu sync.RWMutex
	c  []chan<- Durationer
}

func newTDurationer() *tDurationer {
	return &tDurationer{
		c: make([]chan<- Durationer, 0),
	}
}

func (t *tDurationer) Publish(ctx context.Context, c <-chan Durationer) {
	const op = "topic.tDurationer.publish"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	go func() {
	loop:
		select {
		case <-ctx.Done():
			return
		default:
		}
		for v := range c {
			for _, c := range t.c {
				go func(c chan<- Durationer, v Durationer) {
					select {
					case <-ctx.Done():
						return
					case c <- v:
					}
				}(c, v)
			}
			goto loop
		}
	}()

}

func (t *tDurationer) Publisher(ctx context.Context) chan<- Durationer {
	const op = "topic.tDurationer.publisher"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}

	c := make(chan Durationer)
	t.Publish(ctx, c)
	return c
}

func (t *tDurationer) Subscribe(c chan<- Durationer) {
	const op = "topic.tDurationer.subscribe"

	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	t.mu.Lock()
	t.c = append(t.c, c)
	t.mu.Unlock()
}

type mapDurationer struct {
	mu sync.RWMutex
	m  map[interface{}]*tDurationer
}

func newMapDurationer() *mapDurationer {
	return &mapDurationer{
		m: make(map[interface{}]*tDurationer),
	}
}

func (m *mapDurationer) get(topic interface{}) TopicDurationer {
	const op = "topic.mapDurationer.get"

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
	v = newTDurationer()
	m.m[topic] = v
	return v
}
