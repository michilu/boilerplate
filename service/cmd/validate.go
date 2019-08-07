package cmd

import (
	"os"
	"runtime"

	valid "github.com/asaskevich/govalidator"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
)

func init() {
	valid.TagMap["direxists"] = isDirExists
	valid.TagMap["fileexists"] = isFileExists
	valid.TagMap["filepath"] = isFilePath
}

func isDirExists(s string) bool {
	const op = op + ".isDirExists"

	if s == "" {
		return false
	}
	i, err := os.Stat(s)
	if err != nil {

		slog.Logger().Debug().
			Str("op", op).
			Str("s", s).
			Err(&errs.Error{Op: op, Err: err}).
			Msg("error")

		return false
	}
	if !i.IsDir() {

		slog.Logger().Debug().
			Str("op", op).
			Str("s", s).
			Msgf("%s is a file", s)

		return false
	}

	return true
}

func isFileExists(s string) bool {
	const op = op + ".isFileExists"

	if s == "" {
		return false
	}
	i, err := os.Stat(s)
	if err != nil {

		slog.Logger().Debug().
			Str("op", op).
			Str("s", s).
			Err(&errs.Error{Op: op, Err: err}).
			Msg("error")

		return false
	}
	if i.IsDir() {

		slog.Logger().Debug().
			Str("op", op).
			Str("s", s).
			Msgf("%s is a directory", s)

		return false
	}

	return true
}

func isFilePath(s string) bool {
	ok, os := valid.IsFilePath(s)
	if !ok {
		return false
	}
	switch os {
	case valid.Win:
		if runtime.GOOS != "windows" {
			return false
		}
	}
	return true
}
