package main

import cfg "github.com/michilu/boilerplate/service/config"

var (
	defaults = []cfg.KV{
		{"service.pprof.addr", ":8888"},
		{"service.update.channel", "release"},
		{"service.update.url", "http://localhost:8000/"},
	}
)
