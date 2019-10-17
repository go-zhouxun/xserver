package main

import (
	"fmt"
	"time"

	"github.com/go-zhouxun/xserver/xreq"
	"github.com/go-zhouxun/xserver/xresp"

	"github.com/go-zhouxun/xlog"
	"github.com/go-zhouxun/xserver"
)

func main() {
	logger := xlog.NewDailyLog("log/example.log")
	server := xserver.NewXServer(logger)
	server.Router.Get("/aaa", func(req *xreq.XReq) *xresp.XResp {
		time.Sleep(1 * time.Second)
		req.XContext.LogTime("hhh")
		return xresp.Return("aaa")
	})
	if err := server.Listen(8010); err != nil {
		fmt.Println("listen port :8010 failed, ", err)
	}
}
