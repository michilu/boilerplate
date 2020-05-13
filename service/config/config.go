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
	for _, v := range config {
		v0 := v
		for _, v := range v0 {
			v1 := v
			viper.SetDefault(v1.K, v1.V)
		}
	}
}
