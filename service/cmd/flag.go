package cmd

import (
	"fmt"

	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
)

type opt struct {
	C string `valid:"fileexists"`
}

func LoggingConfig() {
	const op = op + ".LoggingConfig"
	slog.Logger().Info().
		Str("op", op).
		Str("viper", fmt.Sprintf("%v", viper.AllSettings())).
		Msg("config")
}
