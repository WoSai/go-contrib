package logger

import (
	"context"
	"fmt"
	"os"
)

type (
	Level int
)

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
)

const (
	DefaultLevelKey   = "level"
	DefaultMessageKey = "msg"
	DefaultLevel      = Info
)

var (
	levelLogValues = map[Level]string{
		Debug: "debug",
		Info:  "info",
		Warn:  "warn",
		Error: "error",
		Fatal: "fatal",
	}
)

type (
	Helper struct {
		log        Logger
		level      Level
		levelKey   string
		messageKey string
	}

	Option func(*Helper)
)

func NewHelper(logger Logger, opts ...Option) *Helper {
	helper, ok := logger.(*Helper)
	if ok {
		return NewHelper(helper.log, opts...)
	}
	helper = &Helper{log: logger, level: DefaultLevel, levelKey: DefaultLevelKey, messageKey: DefaultMessageKey}
	for _, opt := range opts {
		opt(helper)
	}
	return helper
}

func (h *Helper) Log(keyvals ...interface{}) {
	h.log.Log(keyvals...)
}

func (h *Helper) logWithLevel(level Level, keyvals ...interface{}) {
	if level < h.level {
		return
	}

	kvs := make([]interface{}, 0, 2+len(keyvals))
	kvs = append(kvs, h.levelKey, levelLogValues[h.level])
	kvs = append(kvs, keyvals...)
	h.log.Log(kvs...)

	if level >= Fatal {
		os.Exit(1)
	}
}

func (h *Helper) printfWithLevel(level Level, format string, a ...interface{}) {
	if level < h.level {
		return
	}
	h.log.Log(h.levelKey, levelLogValues[h.level], h.messageKey, fmt.Sprintf(format, a...))
}

func (h *Helper) Debug(keyvals ...interface{}) {
	h.logWithLevel(Debug, keyvals...)
}

func (h *Helper) Debugf(format string, a ...interface{}) {
	h.printfWithLevel(Debug, format, a...)
}

func (h *Helper) Info(keyvals ...interface{}) {
	h.logWithLevel(Info, keyvals...)
}

func (h *Helper) Infof(format string, a ...interface{}) {
	h.printfWithLevel(Info, format, a...)
}

func (h *Helper) Warn(keyvals ...interface{}) {
	h.logWithLevel(Warn, keyvals...)
}
func (h *Helper) Warnf(format string, a ...interface{}) {
	h.printfWithLevel(Warn, format, a...)
}

func (h *Helper) Error(keyvals ...interface{}) {
	h.logWithLevel(Error, keyvals...)
}
func (h *Helper) Errorf(format string, a ...interface{}) {
	h.printfWithLevel(Error, format, a...)
}

func (h *Helper) Fatal(keyvals ...interface{}) {
	h.logWithLevel(Fatal, keyvals...)
}
func (h *Helper) Fatalf(format string, a ...interface{}) {
	h.printfWithLevel(Fatal, format, a...)
}

func (h *Helper) WithContext(ctx context.Context) *Helper {
	return &Helper{
		log:        WithContext(ctx, h.log),
		level:      h.level,
		levelKey:   h.levelKey,
		messageKey: h.messageKey,
	}
}

func WithLevel(level Level) Option {
	return func(h *Helper) {
		h.level = level
	}
}

func WithLevelKey(key string) Option {
	return func(h *Helper) {
		h.levelKey = key
	}
}

func WithMessageKey(key string) Option {
	return func(h *Helper) {
		h.messageKey = key
	}
}
