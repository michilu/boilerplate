package client

import (
	// https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto
	_ "github.com/mwitkow/go-proto-validators"
)

//go:generate interfacer -for github.com/michilu/boilerplate/service/client.Client -as client.Clienter -o entity-Clienter.go
//go:generate interfacer -for github.com/michilu/boilerplate/service/client.Debug -as client.Debuger -o entity-Debuger.go
