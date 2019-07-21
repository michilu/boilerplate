package domain

import (
	"time"

	"github.com/michilu/boilerplate/v/errs"
	"github.com/michilu/boilerplate/v/log"
	"google.golang.org/grpc/codes"
)

// Update ...
//
//go:generate genny -in=../v/pipe/pipe.go -out=gen-pipe-Update.go -pkg=$GOPACKAGE gen "Name=update InT=time.Time OutT=Durationer"
func Update(m time.Time) (Durationer, error) {
	const op = op + ".Update"
	{
		log.Logger().Debug().Str("op", op).Msg("start")
		defer log.Logger().Debug().Str("op", op).Msg("end")
	}
	return nil, nil
}

// Restart ...
//
//go:generate genny -in=../v/pipe/pipe.go -out=gen-pipe-Restart.go -pkg=$GOPACKAGE gen "Name=restart InT=Durationer OutT=struct{}"
func Restart(m Durationer) (struct{}, error) {
	const op = op + ".Restart"
	{
		log.Logger().Debug().Str("op", op).Msg("start")
		defer log.Logger().Debug().Str("op", op).Msg("end")
	}
	if m == nil {
		log.Logger().Warn().Str("op", op).Msg("nil")
		return struct{}{}, nil
	}
	if err := m.Validate(); err != nil {
		const op = op + ".Validate"
		return struct{}{}, &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	return struct{}{}, nil
}
