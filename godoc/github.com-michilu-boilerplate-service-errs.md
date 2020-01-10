# errs
--
    import "github.com/michilu/boilerplate/service/errs"


## Usage

#### type Error

```go
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
```

Error defines a standard application error.

#### func (*Error) Error

```go
func (e *Error) Error() string
```
Error returns the string representation of the error message.

#### func (*Error) MarshalZerologObject

```go
func (p *Error) MarshalZerologObject(e *zerolog.Event)
```
