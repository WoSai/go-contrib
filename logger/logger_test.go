package logger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewStdLogger(os.Stdout)
	logger.Log("hello", "world")
}
