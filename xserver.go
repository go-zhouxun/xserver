package xserver

import (
	"fmt"
	"net/http"

	"github.com/go-zhouxun/xlog"
	"github.com/go-zhouxun/xserver/xreq"
	"github.com/go-zhouxun/xserver/xresp"
	"github.com/go-zhouxun/xserver/xrouter"
	"github.com/go-zhouxun/xutil/xtime"
)

func init() {
	xserver := &XServer{}
	http.HandleFunc("/", xserver.Service)
}

type XServer struct {
	logger xlog.XLog
	router xrouter.XRouter
}

func ParseRequest(req *xreq.XReq) error {
	return req.ParseRequest()
}

func (server XServer) Service(w http.ResponseWriter, r *http.Request) {
	req := xreq.New(r, w)
	_ = req.ParseRequest()
	router := server.router.GetXRouter(req.Path)
	var resp *xresp.XResp
	if router != nil {
		resp = router.Invoke(req)
	} else {
		resp = xresp.NotFound()
	}
	sendResp(req, resp)
	logAccess(server.logger, req, resp)
}

func sendResp(req *xreq.XReq, xresp *xresp.XResp) {
	w := req.W
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-ReqId", req.ReqId)
	w.Header().Set("X-Cost", fmt.Sprintf("%d", xtime.Now()-req.StartTime))
	if len(xresp.Headers) > 0 {
		for key, value := range xresp.Headers {
			w.Header().Set(key, value)
		}
	}
	w.WriteHeader(xresp.HttpCode)
	_, _ = w.Write(xresp.Body)
}

func logAccess(logger xlog.XLog, req *xreq.XReq, xresp *xresp.XResp) {
	// TODO
	logger.Info("")
}
