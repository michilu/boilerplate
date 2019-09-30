package topic

import (
	"context"
	"fmt"

	"github.com/cheekybits/genny/generic"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

// T is a placeholder for the genny.
type T generic.Type

// Ier is a placeholder for the genny.
type Ier generic.Type

// TWithContext is T with context.Context.
type TWithContext struct {
	Context context.Context
	T       Ier
}

// GetContext returns context.Context.
func (p *TWithContext) GetContext() context.Context {
	return p.Context
}

// GetT returns Ier.
func (p *TWithContext) GetT() Ier {
	return p.T
}

// MarshalZerologObject writes TWithContext to given zerolog.Event.
func (p *TWithContext) MarshalZerologObject(e *zerolog.Event) {
	if p.T == nil {
		return
	}
	v, ok := p.T.(zerolog.LogObjectMarshaler)
	if !ok {
		return
	}
	e.Object("TWithContext", v)
}

// String returns TWithContext as string.
func (p *TWithContext) String() string {
	const v0 = "TWithContext<Context: %v, T: %v>"
	if p.T == nil {
		return fmt.Sprintf(v0, p.Context, p.T)
	}
	v1, ok := p.T.(fmt.Stringer)
	if !ok {
		return fmt.Sprintf(v0, p.Context, p.T)
	}
	return fmt.Sprintf(v0, p.Context, v1.String())
}

// Validate returns error if failed validate.
func (p *TWithContext) Validate() error {
	const op = op + ".TWithContext.Validate"
	if p.Context == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*TWithContext.Context' is nil"}
		return err
	}
	if p.T == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*TWithContext.T' is nil"}
		return err
	}
	v0, ok := p.T.(interface{ Validate() error })
	if !ok {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*TWithContext.T' has not 'Validate() error'"}
		return err
	}
	{
		err := v0.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			return err
		}
	}
	return nil
}
