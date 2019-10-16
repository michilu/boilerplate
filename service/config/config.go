package config

import (
	"github.com/spf13/viper"
)

const (
	op = "service/config"
)

type (
	KV struct {
		K string
		V interface{}
	}
)

// SetDefault sets default values to config.
func SetDefault(config ...[]KV) {
	for _, c := range config {
		for _, v := range c {
			viper.SetDefault(v.K, v.V)
		}
	}
}
