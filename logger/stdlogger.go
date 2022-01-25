package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sync"
)

// stdLogger implements Logger by invoking the stdlib log.
type stdLogger struct {
	log  *log.Logger
	pool sync.Pool
}

var _ Logger = (*stdLogger)(nil)

// NewStdLogger new a logger implements Logger interface
func NewStdLogger(output io.Writer) Logger {
	return &stdLogger{
		log: log.New(output, "", log.LstdFlags),
		pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(nil)
			},
		},
	}
}

// Log implementsimplements Logger.Log(...interfaced{}) method
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
