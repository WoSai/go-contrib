package validation

import (
	"strings"
)

type (
	Notification interface {
		Add(error)
		Err() error
	}

	ErrNotification struct {
		errors []string
	}

	notification struct {
		err *ErrNotification
	}
)

func newErrNotification() *ErrNotification {
	return &ErrNotification{errors: []string{}}
}

func (e *ErrNotification) add(msg string) {
	e.errors = append(e.errors, msg)
}

func (e *ErrNotification) errCounts() int {
	return len(e.errors)
}

func (e *ErrNotification) Error() string {
	t := make([]string, 0, 1+len(e.errors))
	t = append(t, "raise errors:")
	t = append(t, e.errors...)
	return strings.Join(t, "\n  - ")
}

func NewSimpleNotification() Notification {
	return &notification{}
}

func (notify *notification) Add(err error) {
	if err == nil {
		return
	}

	if notify.err == nil {
		notify.err = newErrNotification()
	}
	notify.err.add(err.Error())
}

func (notify *notification) Err() error {
	if notify.err == nil || notify.err.errCounts() == 0 {
		return nil
	}
	return notify.err
}
