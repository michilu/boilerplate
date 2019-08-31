package update

import (
	"context"

	"github.com/spf13/cobra"
	"go.opencensus.io/trace"
)

const (
	op = "application/update"
)

func Run(_ *cobra.Command, _ []string) {
	const op = op + ".Run"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)

	Dataflow(ctx)
}
