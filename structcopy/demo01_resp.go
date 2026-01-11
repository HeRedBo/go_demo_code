package structcopy

type Reponse struct {
	Code int         `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Resp interface {
	SetCode(code int) Resp
	SetMsg(msg string) Resp
	Response(errCode interface{}, data interface{}) Reponse
}

type Gin struct {
	response Reponse
}

func (g *Gin) SetCode(code int) Resp {
	g.response.Code = code
	return g
}
func (g *Gin) SetMsg(msg string) Resp {
	g.response.Msg = msg
	return g
}

func (g *Gin) Response(errCode interface{}, data interface{}) Reponse {
	switch errCode.(type) {
	case int:
		intCode := errCode.(int)
		g.response.Code = intCode
		g.response.Data = data
	case string:
		strCode := errCode.(string)
		g.response.Msg = strCode
		g.response.Data = data
	}
	if len(g.response.Msg) == 0 {
		g.response.Msg = "默认字符"
	}
	return g.response
}
