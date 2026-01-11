package structcopy

import (
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestResp(t *testing.T) {
	var appG = Gin{}
	var data = map[string]string{
		"name": "hello",
	}
	dump.P(appG.SetMsg("asd").Response(12312, data))
	dump.P(appG.SetMsg("").Response(2222, data))
}
