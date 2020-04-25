package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/michilu/boilerplate/service/meta"
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
		i, err := strconv.ParseInt(build, 10, 64)
		if err == nil {
			m.Build = time.Unix(i, 0).In(time.UTC)
		}
	}
	if branch != "" || hash != "" || tag != "" {
		m.Vcs = &meta.Vcs{
			Branch: branch,
			Hash:   hash,
			Tag:    tag,
		}
	}
	if err := meta.Set(m); err != nil {
		const op = op + ".meta.Set"
		os.Stderr.WriteString(fmt.Sprintf("op: %s: %s\n", op, err))
	}
}
