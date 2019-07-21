package domain

import (
	_ "github.com/mwitkow/go-proto-validators"
)

//go:generate interfacer -for github.com/michilu/boilerplate/domain.Duration -as domain.Durationer -o vo-Durationer.go
//go:generate genny -in=../v/topic/topic.go -out=gen-topic-Durationer.go -pkg=$GOPACKAGE gen "ChanT=Durationer"
