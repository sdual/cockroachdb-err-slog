package roachslog

import (
	"log/slog"
)

const errAttrKey = "error"

func Err(err error) slog.Attr {
	return slog.Any(errAttrKey, err)
}
