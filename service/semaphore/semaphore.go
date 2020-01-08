package semaphore

import (
	"context"

	sem "github.com/marusama/semaphore"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
)

const (
	op = "service/semaphore"
)

var (
	parallel = 1

	semaphore *sem.Semaphore
)

func Init(parallel int) {
	const op = op + ".Init"
	err := SetParallel(parallel)
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Err(err).Str("op", op).Int("value", parallel).Msg(err.Error())
		return
	}
}

// Semaphore counting resizable semaphore synchronization primitive.
type Semaphore = sem.Semaphore

// SetParallel sets a given number to the parallelism number.
func SetParallel(i int) error {
	const op = op + ".SetParallel"

	if i < 1 {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be 1 or more"}
	}

	parallel = i
	return nil
}

// New initializes a new instance of the Semaphore, specifying the maximum number of concurrent entries.
func New(limit int) sem.Semaphore {
	return sem.New(limit)
}

// Acquire enters the semaphore a specified number of times, blocking only until ctx is done.
func Acquire(ctx context.Context, n int) error {
	return (*get()).Acquire(ctx, n)
}

// TryAcquire acquires the semaphore without blocking.
func TryAcquire(n int) bool {
	return (*get()).TryAcquire(n)
}

// Release exits the semaphore a specified number of times and returns the previous count.
func Release(n int) int {
	return (*get()).Release(n)
}

// GetLimit returns current semaphore limit.
func GetLimit() int {
	return (*get()).GetLimit()
}

// GetCount returns current number of occupied entries in semaphore.
func GetCount() int {
	return (*get()).GetCount()
}

func get() *sem.Semaphore {
	if semaphore != nil {
		return semaphore
	}
	s := sem.New(parallel)
	semaphore = &s
	return semaphore
}
