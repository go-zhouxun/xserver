package xserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-zhouxun/xutil/xstring"

	"github.com/go-zhouxun/xlog"
	"github.com/go-zhouxun/xserver/xreq"
	"github.com/go-zhouxun/xserver/xresp"
	"github.com/go-zhouxun/xserver/xrouter"
	"github.com/go-zhouxun/xutil/xtime"
)

var DefaultXServer *XServer

// single instance
func SetXServer(server *XServer) {
	DefaultXServer = server
}

type XServer struct {
	logger xlog.XLog
	Router *xrouter.XRouter
}

func NewXServer(logger xlog.XLog) *XServer {
	return &XServer{
		logger: logger,
		Router: xrouter.NewXRouter(),
	}
}

func (server XServer) Listen(port int) error {
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func ParseRequest(req *xreq.XReq) error {
	return req.ParseRequest()
}

func (server XServer) Service(w http.ResponseWriter, r *http.Request) {
	req := xreq.New(r, w)
	_ = ParseRequest(req)
	router := server.Router.GetXRouter(req.Path)
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
	TAB := "\t"

	req.XContext.Log("query", req.Query)
	req.XContext.Log("body", req.Param)
	req.XContext.Log("cookie", req.Cookies)
	req.XContext.Log("resp", string(xresp.Body))
	req.XContext.Log("sticker", req.Sticker)
	log := xstring.StringJoin(
		req.Method, TAB,
		strconv.Itoa(xresp.HttpCode), TAB,
		req.Path, TAB,
		strconv.Itoa(int(xtime.Now()-req.StartTime)), TAB,
		req.ReqId, TAB,
		req.ClientIP, TAB,
		req.XContext.String(),
	)
	logger.Info(log)
}
