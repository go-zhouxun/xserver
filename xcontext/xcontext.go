package xcontext

import (
	"encoding/json"

	"github.com/go-zhouxun/xutil/xtime"
)

type XContext struct {
	StartTime int64                  `json:"start_time"`
	Profile   map[string]int64       `json:"profile"`
	LogParam  map[string]interface{} `json:"log_param"`
}

func (context XContext) String() string {
	result, _ := json.Marshal(context)
	return string(result)
}

func NewXContext(startTime int64) *XContext {
	return &XContext{
		StartTime: startTime,
		Profile:   make(map[string]int64),
		LogParam:  make(map[string]interface{}),
	}
}

func (context *XContext) LogTime(op string) {
	context.Profile[op] = xtime.Now() - context.StartTime
}

func (context *XContext) Log(key string, value interface{}) {
	context.LogParam[key] = value
}
