package meta

import (
	_ "github.com/mwitkow/go-proto-validators" // https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto
)

//go:generate interfacer -for github.com/michilu/boilerplate/service/meta.Meta -as meta.Metaer -o vo-Metaer.go
