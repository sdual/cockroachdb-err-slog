# roachslog

[![Go](https://github.com/sdual/roachslog/actions/workflows/go-test.yml/badge.svg)](https://github.com/sdual/roachslog/actions/workflows/go-test.yml)

## Examples

```go
ops := slog.HandlerOptions{
  AddSource: true,
  Level:     slog.LevelInfo,
}

handler := slog.NewJSONHandler(os.Stdout, &ops)
rsHandler := roachslog.NewReachSlogHandler(handler)
slog.SetDefault(slog.New(rsHandler))

```
