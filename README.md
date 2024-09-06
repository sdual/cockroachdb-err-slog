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

	err := doSomething()
	slog.Error("error occurred", roachslog.Err(err))
}

func doSomething() error {
	return errors.New("error")
}
```

```json
{
	"time": "2024-09-06T12:44:15.107483+09:00",
	"level": "ERROR",
	"source": {
		"function": "main.main",
		"file": "/Users/sdual/repos/golang/sample/main.go",
		"line": 22
	},
	"msg": "error occurred",
	"error": "error",
	"stacktrace": "\nmain.doSomething\n\t/Users/sdual/repos/golang/sample/main.go:26\nmain.main\n\t/Users/sdual/repos/golang/sample/main.go:21\nruntime.main\n\t/Users/sdual/go/1.22.4/pkg/mod/golang.org/toolchain@v0.0.1-go1.23.1.darwin-arm64/src/runtime/proc.go:272\nruntime.goexit\n\t/Users/sdual/go/1.22.4/pkg/mod/golang.org/toolchain@v0.0.1-go1.23.1.darwin-arm64/src/runtime/asm_arm64.s:1223"
}
```
