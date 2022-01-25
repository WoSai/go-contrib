package logger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewStdLogger(os.Stdout)
	logger.Log("hello", "world", "bytes", []byte("bytes"), "map", map[string]string{"wosai": "shouqianba"}, "bool", true)
}

func TestMissingValue(t *testing.T) {
	logger := NewStdLogger(os.Stdout)
	logger.Log("hello", "world", "non-value")
	logger.Log()
}
