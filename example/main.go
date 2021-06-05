package main

import (
	"github.com/go-zhouxun/xlog"
	"github.com/go-zhouxun/xserver"
	"github.com/go-zhouxun/xserver/xreq"
	"github.com/go-zhouxun/xserver/xresp"
)

func main() {
	server := xserver.NewXServer()
	server.Router.Get("/aaa", aaa, bbb)

	xlog.Info("Server listening on :8010")
	if err := server.Listen(8010); err != nil {
		xlog.Error("listen port :8010 failed, %v", err)
	}
}

func aaa(req *xreq.XReq) *xresp.XResp {
	return xresp.Return(nil)
}

func bbb(req *xreq.XReq) *xresp.XResp {
	return xresp.Return(nil)
}
