package roachslog

import "log/slog"

type RoachSlogHandler struct {
	handler slog.Handler
}
