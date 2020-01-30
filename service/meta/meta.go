package meta

import (
	"encoding/json"
	"runtime"

	"github.com/jeremywohl/flatten"
	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
	"gopkg.in/yaml.v2"
)

const (
	op = "service/meta"
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

func (m Meta) Flatten() map[string]interface{} {
	var v0 map[string]interface{}
	{
		err := json.Unmarshal(m.JSON(), &v0)
		if err != nil {
			panic(err)
		}
	}
	v1, err := flatten.Flatten(v0, "", flatten.PathStyle)
	if err != nil {
		panic(err)
	}
	return v1
}
