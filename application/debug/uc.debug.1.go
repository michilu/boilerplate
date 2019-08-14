package debug

import (
	fmt "fmt"

	"github.com/google/uuid"
	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
)

func GenerateUUID() (string, error) {
	const op = op + ".GenerateUUID"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	const s = "application.debug.client.id"
	v0 := viper.GetString(s)
	v1, err := uuid.Parse(v0)
	if err == nil {
		return v1.String(), nil
	} else {
		const op = op + ".uuid.Parse"
		slog.Logger().Warn().Str("op", op).Err(err).Str("value", v0).Msg(fmt.Sprintf("check '%s' in config.toml", s))
	}
	v2, err := debug.NewID()
	if err != nil {
		return "", err
	}
	slog.Logger().Warn().Str("op", op).Str("value", v2).Msg(fmt.Sprintf("you can set to '%s' in config.toml", s))
	return v2, nil
}
