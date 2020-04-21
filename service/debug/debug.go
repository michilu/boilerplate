package debug

import (
	"github.com/google/uuid"
	"github.com/michilu/boilerplate/pb"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"google.golang.org/grpc/codes"
)

const (
	op = "service/debug"
)

//go:generate genny -in=../topic/with-context.go -out=gen-ClientWithContext.go -pkg=$GOPACKAGE gen "T=Client Ier=Clienter"
//go:generate interfacer -for github.com/michilu/boilerplate/service/debug.ClientWithContext -as debug.ClientWithContexter -o if-ClientWithContexter.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-Clienter.go            -pkg=$GOPACKAGE gen ChanT=Clienter
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-ClientWithContexter.go -pkg=$GOPACKAGE gen ChanT=ClientWithContexter

type Clienter = pb.DebugClienter

func NewID() (string, error) {
	const op = op + ".NewID"
	{
		slog.Logger().Debug().Str("op", op).Msg(op + ": start")
		defer slog.Logger().Debug().Str("op", op).Msg(op + ": end")
	}
	v0, err := uuid.NewRandom()
	if err != nil {
		const op = op + ".uuid.NewRandom"
		return "", &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
	}
	return v0.String(), nil
}
