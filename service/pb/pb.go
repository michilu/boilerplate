package pb

import (
	"github.com/rs/zerolog"
)

const (
	op = "service/slog/pb"
)

var (
	_ zerolog.LogObjectMarshaler = T{}
)

// T is a placeholder for the genny.
type T struct{}
