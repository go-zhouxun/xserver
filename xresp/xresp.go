package xresp

import (
	"net/http"
)

type XResp struct {
	HttpCode int
	Headers  map[string]string
	Cookies  []http.Cookie
	Body     []byte
}

func Success(result interface{}) *XResp {
	body := &XRespBody{0, "ok", result}
	return &XResp{
		HttpCode: 200,
		Headers:  make(map[string]string),
		Body:     body.ToBytes(),
	}
}

func Fail(status int, msg string, result interface{}) *XResp {
	body := &XRespBody{status, msg, result}
	return &XResp{
		HttpCode: 200,
		Headers:  make(map[string]string),
		Body:     body.ToBytes(),
	}
}

func Return(result interface{}) *XResp {
	body := &XRespBody{0, "OK", result}
	return &XResp{
		HttpCode: 200,
		Headers:  make(map[string]string),
		Body:     body.ToBytes(),
	}
}

func NotFound() *XResp {
	resp := &XResp{
		HttpCode: 404,
		Body:     []byte("{\"status\":-1, \"msg\": \"not found\"}"),
	}
	return resp
}
