package config

import (
	"github.com/spf13/viper"
)

type (
	kv struct {
		k string
		v interface{}
	}
)

var (
	config = []kv{
		{"service.update.channel", "release"},
		{"service.update.url", "http://localhost:8000/"},
	}
)

// SetDefault sets default values to config.
func SetDefault() {
	for _, c := range [][]kv{config} {
		for _, v := range c {
			viper.SetDefault(v.k, v.v)
		}
	}
}
