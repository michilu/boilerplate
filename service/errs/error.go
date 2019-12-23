package errs

// Ref: https://middlemost.com/failure-is-your-domain/

import (
	"fmt"
	"strings"
)

const (
	op = "servic/errs"
)

// Error defines a standard application error.
type Error struct {
	// Code is a stringable type as defined in the gRPC spec.
	// https://godoc.org/google.golang.org/grpc/codes#Code
	Code fmt.Stringer

	// Message is a human-readable explanation specific to this occurrence of the error.
	Message string

	// Logical operation and nested error.
	Op  string
	Err error

	errorMessage string
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	if e.errorMessage != "" {
		return e.errorMessage
	}

	const c0 = ": "
	var (
		v0, v1, v2, v3 string
		v4             strings.Builder
	)

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		v0 = e.Err.Error()
	} else {
		if e.Code != nil {
			v1 = e.Code.String()
		}
		v3 = e.Message
		if v3 != "" {
			v2 = c0
		}
	}

	v4.Grow(len(e.Op) + len(c0) + len(v0) + len(v1) + len(v2) + len(v3))

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprint(&v4, e.Op, c0)
	}

	fmt.Fprint(&v4, v0, v1, v2, v3)
	e.errorMessage = v4.String()
	return e.errorMessage
}
