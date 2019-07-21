package meta

import (
	"runtime"

	"github.com/michilu/boilerplate/v/errs"
	"github.com/michilu/boilerplate/v/log"
	"google.golang.org/grpc/codes"
	"gopkg.in/yaml.v2"
)

const (
	op = "meta"

	BuildFmt = "Jan 2 15:04:05 2006"
)

var (
	m Meta
)

// Set sets a meta data.
func Set(v *Meta) {
	const op = op + ".Set"
	v.Runtime = &Runtime{
		Arch:    runtime.GOARCH,
		Os:      runtime.GOOS,
		Version: runtime.Version(),
	}
	if err := v.Validate(); err != nil {
		const op = op + ".Validate"
		log.Logger().Error().Str("op", op).Err(&errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}).Msg("error")
	}
	m = *v
}

// Yaml returns Meta as YAML.
func Yaml() (string, error) {
	const op = op + ".Yaml"
	b, err := yaml.Marshal(&m)
	if err != nil {
		const op = op + ".yaml.Marshal"
		return "", &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	return string(b), nil
}

// Name returns a name.
func Name() string {
	return m.GetName()
}
