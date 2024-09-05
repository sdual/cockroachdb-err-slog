# roachslog

[![Go](https://github.com/sdual/roachslog/actions/workflows/go-test.yml/badge.svg)](https://github.com/sdual/roachslog/actions/workflows/go-test.yml)

## Examples

```go
package main

import (
	"log/slog"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/sdual/roachslog"
)

func main() {
	ops := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	handler := slog.NewJSONHandler(os.Stdout, &ops)
	rsHandler := roachslog.NewReachSlogHandler(handler)
	slog.SetDefault(slog.New(rsHandler))

	err := errors.New("error")
	slog.Error("error occurred", roachslog.Err(err))
}
```
