package meta

import (
	"fmt"
	"runtime"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/michilu/boilerplate/v/errs"
)

const (
	buildFmt = "Jan 2 15:04:05 2006"
)

var (
	m *meta
)

type (
	Meta struct {
		Build  string
		Hash   string
		Name   string
		SemVer string
		Serial string
	}

	meta struct {
		Build   time.Time `yaml:",omitempty"`
		Hash    string    `yaml:",omitempty"`
		Name    string    `yaml:",omitempty"`
		SemVer  string    `yaml:",omitempty"`
		Serial  string    `yaml:",omitempty"`
		Runtime *runTime  `yaml:",omitempty"`
	}

	runTime struct {
		Version string `yaml:",omitempty"`
		Arch    string `yaml:",omitempty"`
		Os      string `yaml:",omitempty"`
	}
)

func (m meta) String() string {
	o, err := yaml.Marshal(&m)
	if err != nil {
		panic(err)
	}
	return string(o)
}

// Set sets a meta data.
func Set(v *Meta) error {
	const op = "meta.Set"

	m = &meta{
		Name:   v.Name,
		Hash:   v.Hash,
		SemVer: v.SemVer,
		Serial: v.Serial,
		Runtime: &runTime{
			Version: runtime.Version(),
			Arch:    runtime.GOARCH,
			Os:      runtime.GOOS,
		},
	}

	if v.Build == "" {
		return nil
	}
	t, err := time.Parse(buildFmt, v.Build)
	if err != nil {
		return &errs.Error{Op: op, Err: err}
	}
	m.Build = t

	return nil
}

// Get returns a fmt.Stringer.
func Get() fmt.Stringer {
	return *m
}

// Name returns a name.
func Name() string {
	if m == nil {
		return ""
	}
	return m.Name
}
