package xrouter

import (
	"net/http"

	"github.com/go-zhouxun/xserver/xerr"
	"github.com/go-zhouxun/xserver/xresp"

	"github.com/go-zhouxun/xserver/xreq"
)

const (
	GET  = 1 << iota
	POST = 1 << iota
	ALL  = GET | POST
)

type XHandler func(*xreq.XReq) *xresp.XResp

type Info struct {
	HttpMethod int32 //GET POST PUT OPTION DELETE ...
	Handlers   []XHandler
}

type XRouter struct {
	mapping map[string]*Info
}

func NewXRouter() *XRouter {
	return &XRouter{mapping: make(map[string]*Info)}
}

func (router *XRouter) GetXRouter(url string) *Info {
	return router.mapping[url]
}

func (routeInfo *Info) Invoke(req *xreq.XReq) *xresp.XResp {
	defaultResp := &xresp.XResp{
		HttpCode: http.StatusInternalServerError,
		Body:     []byte(`{"status": -1, "msg": "未知错误"}`),
	}
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(xerr.XErr)
			if ok {
				defaultResp.HttpCode = http.StatusBadRequest
				defaultResp.Body = []byte(err.Error())
			}
		}
	}()
	for _, handler := range routeInfo.Handlers {
		if resp := handler(req); resp != nil {
			return resp
		}
	}
	return defaultResp
}

func (router *XRouter) Get(url string, handlers ...XHandler) {
	routerInfo := &Info{
		HttpMethod: GET,
		Handlers:   handlers,
	}
	router.mapping[url] = routerInfo
}

func (router *XRouter) Post(url string, handlers ...XHandler) {
	routeInfo := &Info{
		HttpMethod: POST,
		Handlers:   handlers,
	}
	router.mapping[url] = routeInfo
}

func (router *XRouter) All(url string, handlers ...XHandler) {
	routeInfo := &Info{
		HttpMethod: ALL,
		Handlers:   handlers,
	}
	router.mapping[url] = routeInfo
}
