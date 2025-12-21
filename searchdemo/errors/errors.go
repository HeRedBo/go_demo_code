package errors

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"runtime"
)

func callers() []uintptr {
	var pcs [32]uintptr
	l := runtime.Callers(3, pcs[:])
	return pcs[:l]
}

type Error interface {
	error
}

type item struct {
	msg   string
	stack []uintptr
}

func (i *item) Error() string {
	return i.msg
}

func (i *item) Format(s fmt.State, verb rune) {
	io.WriteString(s, i.msg)
	io.WriteString(s, "\n")

	for _, pc := range i.stack {
		fmt.Fprintf(s, "%+v\n", errors.Frame(pc))
	}
}

func New(msg string) Error {
	return &item{msg: msg, stack: callers()}
}

func ErrorF(format string, args ...interface{}) Error {
	return &item{msg: fmt.Sprintf(format, args...), stack: callers()}
}

// Wrap with some extra message into err
func Wrap(err error, msg string) Error {
	if err == nil {
		return nil
	}
	var e *item
	//ok := errors.As(err, &e) Type assertion on errors fails on wrapped errors
	e, ok := err.(*item)
	if !ok {
		return &item{msg: fmt.Sprintf("%s;%s", msg, err.Error()), stack: callers()}
	}

	e.msg = fmt.Sprintf("%s; %s", msg, e.msg)
	return e
}

// Wrapf with some extra message into err
func Wrapf(err error, format string, args ...interface{}) Error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, args...)
	var e *item
	ok := errors.As(err, &e)
	if !ok {
		return &item{msg: fmt.Sprintf("%s;%s", msg, err.Error()), stack: callers()}
	}
	e.msg = fmt.Sprintf("%s; %s", msg, e.msg)
	return e
}

func WithStack(err error) Error {
	if err == nil {
		return nil
	}

	var e *item
	if errors.As(err, &e) {
		return e
	}

	return &item{msg: err.Error(), stack: callers()}

}
