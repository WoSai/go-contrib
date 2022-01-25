package logger

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
)

type (
	Logger interface {
		Log(keyvalues ...interface{})
	}

	stdLogger struct {
		log  *log.Logger
		pool sync.Pool
	}
)

var (
	// ErrMissingValue is appended to keyvalues slices with odd length to substitute the missing value.
	ErrMissingValue = errors.New("(missing value)")

	_ Logger = (*stdLogger)(nil)
)

func NewStdLogger(output io.Writer) *stdLogger {
	return &stdLogger{
		log: log.New(output, "", log.LstdFlags),
		pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(nil)
			},
		},
	}
}

func (logger *stdLogger) Log(keyvalues ...interface{}) {
	if len(keyvalues) == 0 {
		return
	}
	if len(keyvalues)&1 == 1 {
		keyvalues = append(keyvalues, ErrMissingValue.Error())
	}

	buffer := logger.pool.Get().(*bytes.Buffer)
	defer logger.pool.Put(buffer)
	defer buffer.Reset()

	for i := 0; i < len(keyvalues); i += 2 {
		_, _ = fmt.Fprintf(buffer, " %s=%v", keyvalues[i], keyvalues[i+1])
	}

	_ = logger.log.Output(4, buffer.String())
}
