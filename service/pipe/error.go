package pipe

import (
	"github.com/michilu/boilerplate/service/slog"
)

// DefaultErrorHandler is a default error handler.
func DefaultErrorHandler(err error) bool {
	const op = op + ".DefaultErrorHandler"
	slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
	return true
}
