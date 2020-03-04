package topic

import (
	"context"
	"sync"

	"github.com/cheekybits/genny/generic"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
)

var (
	topicChanT *mapChanT
)

func init() {
	topicChanT = newMapChanT()
}

// GetTopicChanT returns a TopicChanT of the given topic.
func GetTopicChanT(topic interface{}) TopicChanT {
	return topicChanT.get(topic)
}

// TopicChanT is a topic.
type TopicChanT interface {
	// Publish returns a '<-chan ChanT' that joins to the given topic.
	Publish(ctx context.Context, c <-chan ChanT)
	// Publisher returns a 'chan<- ChanT' that joins to the given topic.
	Publisher(ctx context.Context) chan<- ChanT
	// Subscribe returns a 'chan<- ChanT' that joins to the given topic.
	Subscribe(c chan<- ChanT)
}

// ChanT is a placeholder for the genny.
type ChanT generic.Type

type tChanT struct {
	mu sync.RWMutex
	c  []chan<- ChanT
}

func newTChanT() *tChanT {
	return &tChanT{
		c: make([]chan<- ChanT, 0),
	}
}

func (t *tChanT) Publish(ctx context.Context, c <-chan ChanT) {
	const op = op + ".tChanT.Publish"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	go slog.Recover(ctx, func(ctx context.Context) error {
	loop:
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		for v := range c {
			for _, c := range t.c {
				go slog.Recover(ctx, func(ctx context.Context) error {
					func(c chan<- ChanT, v ChanT) {
						select {
						case <-ctx.Done():
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

func (t *tChanT) Publisher(ctx context.Context) chan<- ChanT {
	const op = op + ".tChanT.Publisher"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}

	c := make(chan ChanT)
	t.Publish(ctx, c)
	return c
}

func (t *tChanT) Subscribe(c chan<- ChanT) {
	const op = op + ".tChanT.Subscribe"

	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	t.mu.Lock()
	t.c = append(t.c, c)
	t.mu.Unlock()
}

type mapChanT struct {
	mu sync.RWMutex
	m  map[interface{}]*tChanT
}

func newMapChanT() *mapChanT {
	return &mapChanT{
		m: make(map[interface{}]*tChanT),
	}
}

func (m *mapChanT) get(topic interface{}) TopicChanT {
	const op = op + ".mapChanT.get"

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
	v = newTChanT()
	m.m[topic] = v
	return v
}
