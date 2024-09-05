package roachslog

import (
	"context"
	"log/slog"

	"github.com/cockroachdb/errors"
)

const (
	defaultErrorAttrKey     = "error"
	defaultStactraceAttrKey = "stacktrace"
)

type (
	// RoachSlogHandler is the slog handler to format stacktrace of logs cockroachdb/errors.
	RoachSlogHandler struct {
		handler       slog.Handler
		errAttrKey    string
		stacktraceKey string
	}

	opt func(*RoachSlogHandler)
)

// NewReachSlogHandler creates a new RoachSlogHandler instance.
func NewReachSlogHandler(handler slog.Handler, opts ...opt) RoachSlogHandler {
	rs := &RoachSlogHandler{
		handler:       handler,
		errAttrKey:    defaultErrorAttrKey,
		stacktraceKey: defaultStactraceAttrKey,
	}

	for _, option := range opts {
		option(rs)
	}
	return *rs
}

func WithErrorAttrKey(key string) opt {
	return func(rs *RoachSlogHandler) {
		rs.errAttrKey = key
	}
}

func WithStacktraceAttrKey(key string) opt {
	return func(rs *RoachSlogHandler) {
		rs.stacktraceKey = key
	}
}

func (rs *RoachSlogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return rs.handler.Enabled(ctx, level)
}

func (rs *RoachSlogHandler) Handle(ctx context.Context, record slog.Record) error {
	var stacktrace string
	record.Attrs(func(attr slog.Attr) bool {
		if attr.Key == rs.errAttrKey {
			err, ok := attr.Value.Any().(error)
			// If the cast to type error fails, no information about the error is logged.
			// The value of the error data type should be available as the value corresponding to the specified key (errAttrKey).
			if !ok {
				return false
			}
			stacktrace, ok = extractStacktrace(err)
			if !ok {
				return false
			}
		}
		return true
	})

	record.AddAttrs(slog.String(rs.stacktraceKey, stacktrace))
	return rs.handler.Handle(ctx, record)
}

func (rs *RoachSlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &RoachSlogHandler{handler: rs.handler.WithAttrs(attrs)}
}
func (rs *RoachSlogHandler) WithGroup(g string) slog.Handler {
	return &RoachSlogHandler{handler: rs.handler.WithGroup(g)}
}

func extractStacktrace(err error) (string, bool) {
	safeDetails := errors.GetSafeDetails(err).SafeDetails
	if len(safeDetails) > 0 {
		return safeDetails[0], true
	}
	return "", false
}
