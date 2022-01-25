package logger

import (
	"context"
	"os"
	"testing"
)

var (
	testKey = new(struct{})
)

func testValuer() Valuer {
	return func(ctx context.Context) interface{} {
		return ctx.Value(testKey)
	}
}

func TestWith(t *testing.T) {
	logger := NewStdLogger(os.Stdout)
	logger.Log("message", "hello world")

	logger = With(logger, "level", "info")
	logger.Log("message", "hello again")

	ctx := context.WithValue(context.Background(), testKey, "world peace")
	logger = WithContext(ctx, logger)
	logger = With(logger, "value-from-ctx", testValuer(), "caller", DefaultCaller)
	logger.Log("???", "!!!")
}
