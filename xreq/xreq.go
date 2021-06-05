package xreq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-playground/validator"

	"github.com/go-zhouxun/xserver/xcontext"
	"github.com/go-zhouxun/xserver/xerr"
	"github.com/go-zhouxun/xserver/xtype"
	"github.com/go-zhouxun/xutil/xtime"
)

var validate = validator.New()

type XReq struct {
	R         *http.Request
	W         http.ResponseWriter
	Method    string
	Path      string
	ReqId     string
	ClientIP  string
	StartTime int64
	Query     map[string]interface{}
	Param     map[string]interface{}
	Cookies   map[string]string
	Sticker   map[string]interface{}
	Body      []byte
	XContext  *xcontext.XContext
}

func New(r *http.Request, w http.ResponseWriter) *XReq {
	startTime := xtime.Now()
	return &XReq{
		R:         r,
		W:         w,
		Method:    r.Method,
		Path:      r.URL.Path,
		ReqId:     createReqId(),
		ClientIP:  getRealIP(r),
		StartTime: startTime,
		Body:      make([]byte, 0),
		Cookies:   make(map[string]string),
		Query:     make(map[string]interface{}),
		Param:     make(map[string]interface{}),
		Sticker:   make(map[string]interface{}),
		XContext:  xcontext.NewXContext(startTime),
	}
}

func (req *XReq) ParseBody(s interface{}) error {
	err := json.Unmarshal(req.Body, &s)
	if err != nil {
		panic(xerr.NewXErr(-7, err.Error(), nil))
	}
	err = validate.Struct(s)
	if err != nil {
		panic(xerr.NewXErr(-7, err.Error(), nil))
	}
	return nil
}

func (req *XReq) ParseRequest() error {
	req.parseCookie()
	req.parseQuery()
	if strings.ToUpper(req.Method) == "POST" {
		if err := req.parseBody(); err != nil {
			return err
		}
	}
	return nil
}

func (req *XReq) parseCookie() {
	if cookies := req.R.Cookies(); len(cookies) != 0 {
		for _, cookie := range cookies {
			req.Cookies[cookie.Name] = cookie.Value
		}
	}
}

func (req *XReq) parseQuery() {
	query := req.R.RequestURI[strings.LastIndex(req.R.RequestURI, "?")+1:]
	req.Query = kv2map(query)
}

func (req *XReq) parseBody() error {
	body, err := ioutil.ReadAll(req.R.Body)
	if req.R.Body != nil {
		_ = req.R.Body.Close()
	}
	req.Body = body
	if err != nil {
		return xerr.NewXErr(-1, err.Error(), nil)
	}
	contentType := req.R.Header.Get("Content-Type")
	if strings.Contains(contentType, "json") {
		err = json.Unmarshal(body, &req.Param)
		if err != nil {
			return xerr.NewXErr(-2, err.Error(), nil)
		}
	} else if contentType == "application/x-www-from-urlencoded" {
		bodyStr := string(body)
		req.Param = kv2map(bodyStr)
	} else {
		return xerr.NewXErr(-3, "unsupport content-type", nil)
	}
	return nil
}

func kv2map(kv string) map[string]interface{} {
	pairs := strings.Split(kv, "&")
	result := make(map[string]interface{})
	if len(pairs) > 0 {
		for _, pair := range pairs {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 {
				result[kv[0]] = kv[1]
			}
		}
	}
	return result
}

func (req *XReq) GetParam(name string) (interface{}, bool) {
	v, exist := req.Param[name]
	if !exist {
		v, exist = req.Query[name]
	}
	return v, exist
}

func (req *XReq) MustGetInt64(name string) int64 {
	v, exist := req.GetParam(name)
	if !exist {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	value, ok := xtype.GetInt64(v)
	if !ok {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	return value
}

func (req *XReq) MustGetInt32(name string) int32 {
	v, exist := req.GetParam(name)
	if !exist {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	value, ok := xtype.GetInt32(v)
	if !ok {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	return value
}

func (req *XReq) MustGetFloat64(name string) float64 {
	v, exist := req.GetParam(name)
	if !exist {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	value, ok := xtype.GetFloat64(v)
	if !ok {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	return value
}

func (req *XReq) MustGetFloat32(name string) float64 {
	v, exist := req.GetParam(name)
	if !exist {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	value, ok := xtype.GetFloat32(v)
	if !ok {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	return value
}

func (req *XReq) MustGetBool(name string) bool {
	v, exist := req.GetParam(name)
	if !exist {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	value, ok := xtype.GetBool(v)
	if !ok {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	return value
}

func (req *XReq) MustGetString(name string) string {
	v, exist := req.GetParam(name)
	if !exist {
		panic(xerr.NewXErr(-7, "param "+name+" not exist", nil))
	}
	return xtype.V2String(v)
}
