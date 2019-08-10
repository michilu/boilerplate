package meta

import (
	// https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto
	_ "github.com/mwitkow/go-proto-validators"
)

//go:generate interfacer -for github.com/michilu/boilerplate/service/meta.Meta -as meta.Metaer -o vo-Metaer.go
