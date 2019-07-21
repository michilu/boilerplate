package main

import (
	"time"

	"github.com/michilu/boilerplate/v/cmd"
	"github.com/michilu/boilerplate/v/meta"
)

// override by ldflags.
var (
	branch string
	build  string
	hash   string
	serial string
)

func init() {
	const op = "main.init"
	m := &meta.Meta{
		Name:   name,
		Semver: semver,
		Serial: serial,
	}
	if build != "" {
		t, err := time.Parse(meta.BuildFmt, build)
		if err == nil {
			m.Build = t
		}
	}
	if branch != "" || hash != "" {
		m.Vcs = &meta.Vcs{
			Branch: branch,
			Hash:   hash,
		}
	}
	meta.Set(m)
	cmd.Init(ns)
}
