package cmd

import (
	"os"
	"runtime"

	valid "github.com/asaskevich/govalidator"

	"github.com/michilu/boilerplate/v/errs"
	"github.com/michilu/boilerplate/v/log"
)

func init() {
	valid.TagMap["direxists"] = isDirExists
	valid.TagMap["fileexists"] = isFileExists
	valid.TagMap["filepath"] = isFilePath
}

func isDirExists(s string) bool {
	const op = "cmd.validate.isDirExists"

	if s == "" {
		return false
	}
	i, err := os.Stat(s)
	if err != nil {

		log.Logger().Debug().
			Str("op", op).
			Str("s", s).
			Err(&errs.Error{Op: op, Err: err}).
			Msg("error")

		return false
	}
	if !i.IsDir() {

		log.Logger().Debug().
			Str("op", op).
			Str("s", s).
			Msgf("%s is a file", s)

		return false
	}

	return true
}

func isFileExists(s string) bool {
	const op = "cmd.validate.isFileExists"

	if s == "" {
		return false
	}
	i, err := os.Stat(s)
	if err != nil {

		log.Logger().Debug().
			Str("op", op).
			Str("s", s).
			Err(&errs.Error{Op: op, Err: err}).
			Msg("error")

		return false
	}
	if i.IsDir() {

		log.Logger().Debug().
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
