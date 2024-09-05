# roachslog

[![Go](https://github.com/sdual/roachslog/actions/workflows/go-test.yml/badge.svg)](https://github.com/sdual/roachslog/actions/workflows/go-test.yml)

## Examples

```go
	ops := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			switch attr.Key {
			case slog.LevelKey:
				attr = slog.Attr{
					Key:   "severity",
					Value: attr.Value,
				}
			case slog.MessageKey:
				attr = slog.Attr{
					Key:   "message",
					Value: attr.Value,
				}
			case slog.SourceKey:
				attr = slog.Attr{
					Key:   "logging.googleapis.com/sourceLocation",
					Value: attr.Value,
				}
			}
			return attr
		},
	}
	handler := slog.NewJSONHandler(os.Stdout, &ops)
	rsHandler := roachslog.NewReachSlogHandler(handler)
	slog.SetDefault(slog.New(rsHandler))

```
