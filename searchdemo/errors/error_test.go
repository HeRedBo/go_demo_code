package errors

import (
	"errors"
	"fmt"
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("origin error")
	fmt.Println(err)
	dump.P(err)
	err = Wrap(err, "add error")
	fmt.Println(WithStack(err))
	dump.P(WithStack(err))
}
