package main

import (
	"fmt"

	"github.com/go-zhouxun/xlog"
	"github.com/go-zhouxun/xserver"
	"github.com/go-zhouxun/xserver/xreq"
	"github.com/go-zhouxun/xserver/xresp"
)

func main() {
	logger := xlog.NewDailyLog("log/example.log")
	server := xserver.NewXServer(logger)
	server.Router.Get("/aaa", aaa, bbb)
	if err := server.Listen(8010); err != nil {
		fmt.Println("listen port :8010 failed, ", err)
	}
}

func aaa(req *xreq.XReq) *xresp.XResp {
	return xresp.Return(nil)
}

func bbb(req *xreq.XReq) *xresp.XResp {
	return xresp.Return(nil)
}
