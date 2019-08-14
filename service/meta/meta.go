package meta

import (
	"encoding/json"
	"runtime"

	"google.golang.org/grpc/codes"
	"gopkg.in/yaml.v2"

	"github.com/michilu/boilerplate/service/errs"
)

const (
	op = "service/meta"

	BuildFmt = "Jan 2 15:04:05 2006"
)

var (
	m Meta
)

// Get returns a Meta.
func Get() Meta {
	return m
}

// Set sets a meta data.
func Set(v *Meta) error {
	const op = op + ".Set"
	v.Runtime = &Runtime{
		Arch:    runtime.GOARCH,
		Os:      runtime.GOOS,
		Version: runtime.Version(),
	}
	m = *v
	if err := v.Validate(); err != nil {
		const op = op + ".Validate"
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	return nil
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

func (m Meta) JSON() []byte {
	b, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}
	return b
}
