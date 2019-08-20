package main

import (
	"time"

	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/slog"
)

const (
	name = "boilerplate"
)

// override by ldflags.
var (
	semver  = "0.0.0.0-alpha"
	branch  string
	build   string
	channel string
	hash    string
	serial  string
	tag     string
)

func init() {
	const op = op + ".init"
	m := &meta.Meta{
		Name:    name,
		Semver:  semver,
		Channel: channel,
		Serial:  serial,
	}
	if build != "" {
		t, err := time.Parse(meta.BuildFmt, build)
		if err == nil {
			m.Build = t
		}
	}
	if branch != "" || hash != "" || tag != "" {
		m.Vcs = &meta.Vcs{
			Branch: branch,
			Hash:   hash,
			Tag:    tag,
		}
	}
	err := meta.Set(m)
	if err != nil {
		const op = op + ".meta.Set"
		slog.Logger().Error().Str("op", op).Err(err).Msg("error")
	}
}
