package pb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
)

const (
	op = "pb"
)

// debug.proto
//go:generate genny -in=../service/pb/zerolog.go -out=gen-zerolog-DebugClient.go gen T=DebugClient
//go:generate interfacer -for github.com/michilu/boilerplate/pb.DebugClient -as pb.DebugClienter -o if-DebugClienter.go

// event.proto
//go:generate genny -in=../service/pb/zerolog.go -out=gen-zerolog-Event.go gen T=Event
//go:generate interfacer -for github.com/michilu/boilerplate/pb.Event -as pb.Eventer -o if-Eventer.go

// keyvalue.proto
//go:generate genny -in=../service/pb/zerolog.go -out=gen-zerolog-Key.go       gen T=Key
//go:generate genny -in=../service/pb/zerolog.go -out=gen-zerolog-KeyValue.go  gen T=KeyValue
//go:generate genny -in=../service/pb/zerolog.go -out=gen-zerolog-KeyPrefix.go gen T=KeyPrefix
//go:generate interfacer -for github.com/michilu/boilerplate/pb.Key       -as pb.Keyer       -o if-Keyer.go
//go:generate interfacer -for github.com/michilu/boilerplate/pb.KeyValue  -as pb.KeyValuer   -o if-KeyValuer.go
//go:generate interfacer -for github.com/michilu/boilerplate/pb.KeyPrefix -as pb.KeyPrefixer -o if-KeyPrefixer.go

// type.proto
//go:generate genny -in=../service/pb/zerolog.go -out=gen-zerolog-Bool.go gen T=Bool
//go:generate genny -in=../service/pb/zerolog.go -out=gen-zerolog-Byte.go gen T=Byte
