package logger

import "context"

type (
	Valuer func(context.Context) interface{}
)
