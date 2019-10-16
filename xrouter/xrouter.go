package xrouter

import (
	"sync"

	"github.com/go-zhouxun/xserver/xresp"

	"github.com/go-zhouxun/xserver/xreq"
)

const (
	GET  = 1 << iota
	POST = 1 << iota
	ALL  = GET | POST
)

type XHandler func(*xreq.XReq) *xresp.XResp

type XRouterInfo struct {
	HttpMethod int32 //GET POST PUT OPTION DELETE ...
	Handlers   []XHandler
}

type XRouter struct {
	mapping map[string]*XRouterInfo
	Lock    sync.Locker
}

var Router = &XRouter{
	mapping: make(map[string]*XRouterInfo),
	Lock:    new(sync.RWMutex),
}

func (xRouter *XRouter) GetXRouter(url string) *XRouterInfo {
	return xRouter.mapping[url]
}

func (routeInfo *XRouterInfo) Invoke(req *xreq.XReq) *xresp.XResp {
	for _, handler := range routeInfo.Handlers {
		if resp := handler(req); resp != nil {
			return resp
		}
	}
	return nil
}

func (router *XRouter) Get(url string, handlers ...XHandler) {
	Router.Lock.Lock()
	defer Router.Lock.Unlock()
	routerInfo := &XRouterInfo{
		HttpMethod: GET,
		Handlers:   handlers,
	}
	Router.mapping[url] = routerInfo
}

func (router *XRouter) Post(url string, handlers ...XHandler) {
	Router.Lock.Lock()
	defer Router.Lock.Unlock()
	routeInfo := &XRouterInfo{
		HttpMethod: POST,
		Handlers:   handlers,
	}
	Router.mapping[url] = routeInfo
}

func (router *XRouter) All(url string, handlers ...XHandler) {
	Router.Lock.Lock()
	defer Router.Lock.Unlock()
	routeInfo := &XRouterInfo{
		HttpMethod: ALL,
		Handlers:   handlers,
	}
	Router.mapping[url] = routeInfo
}
