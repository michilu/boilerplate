package meta

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
)

//go:generate interfacer -for github.com/michilu/boilerplate/service/meta.Meta -as meta.Metaer -o vo-Metaer.go
