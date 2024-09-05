package roachslog

import (
	"context"
	"log/slog"

	"github.com/cockroachdb/errors"
)

const defaultStactraceAttrKey = "stacktrace"

type (
	// RoachSlogHandler is the slog handler to format stacktrace of logs cockroachdb/errors.
	RoachSlogHandler struct {
		handler       slog.Handler
		stacktraceKey string
	}

	opt func(*RoachSlogHandler)
)

// NewReachSlogHandler creates a new RoachSlogHandler instance.
func NewReachSlogHandler(handler slog.Handler, opts ...opt) slog.Handler {
	rs := &RoachSlogHandler{
		handler:       handler,
		stacktraceKey: defaultStactraceAttrKey,
	}

	for _, option := range opts {
		option(rs)
	}
	return rs
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
		if attr.Key == errAttrKey {
			// If the cast to error data type fails, no information about the error is logged.
			// The value of the error data type should be available as the value corresponding to the specified key (errAttrKey).
			if err, ok := attr.Value.Any().(error); ok {
				stacktrace = extractStacktrace(err)
				record.AddAttrs(slog.String(rs.stacktraceKey, stacktrace))
			}
			return false
		}
		return true
	})

	return rs.handler.Handle(ctx, record)
}

func (rs *RoachSlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &RoachSlogHandler{handler: rs.handler.WithAttrs(attrs)}
}
func (rs *RoachSlogHandler) WithGroup(g string) slog.Handler {
	return &RoachSlogHandler{handler: rs.handler.WithGroup(g)}
}

func extractStacktrace(err error) string {
	safeDetails := errors.GetSafeDetails(err).SafeDetails
	if len(safeDetails) > 0 {
		return safeDetails[0]
	}
	return ""
}
