package xerr

import "encoding/json"

type XErr struct {
	Code    int64
	Msg     string
	Sticker map[string]interface{}
}

func NewXErr(code int64, msg string, sticker map[string]interface{}) XErr {
	return XErr{
		Code:    code,
		Msg:     msg,
		Sticker: sticker,
	}
}

func (err XErr) Error() string {
	detail, _ := json.Marshal(err)
	return string(detail)
}
