package logger

import (
	"context"
	"os"
	"testing"
)

func TestHelper(t *testing.T) {
	logger := NewStdLogger(os.Stdout)
	helper := NewHelper(logger, WithMessageKey("message"))
	helper.Info("hello", "world")
	helper.Infof("%s", "foobar!")

	helper = helper.WithContext(context.TODO())
	helper.Info("???", "!!!!")
}
