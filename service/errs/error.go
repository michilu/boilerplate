package errs

// Ref: https://middlemost.com/failure-is-your-domain/

import (
	"bytes"
	"fmt"
)

const (
	op = "servic.errs"
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
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	const op = op + ".Error.Error()"

	// `bytes.Buffer.WriteString` always returns a nil error.
	// https://golang.org/pkg/bytes/#Buffer.WriteString
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		buf.WriteString(e.Op + ": ") // #nosec
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error()) // #nosec
	} else {
		if e.Code != nil {
			buf.WriteString(e.Code.String() + " ") // #nosec
		}
		buf.WriteString(e.Message) // #nosec
	}

	return buf.String()
}
