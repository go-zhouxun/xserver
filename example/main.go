package main

import (
	"os"

	"github.com/go-zhouxun/xlog"
	"github.com/go-zhouxun/xserver"
	"github.com/go-zhouxun/xserver/xreq"
	"github.com/go-zhouxun/xserver/xresp"
)

func main() {
	server := xserver.NewXServer()
	server.Router.Get("/aaa", aaa, bbb)

	port := os.Getenv("APP_PORT")
	xlog.Info("Server listening on %s", port)
	if err := server.Listen(port); err != nil {
		xlog.Error("listen port %s failed, %v", port, err)
	}
}

func aaa(req *xreq.XReq) *xresp.XResp {
	return xresp.Return(nil)
}

func bbb(req *xreq.XReq) *xresp.XResp {
	return xresp.Return(nil)
}
