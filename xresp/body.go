package xresp

import "encoding/json"

type XRespBody struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func (xRespBody *XRespBody) ToBytes() (bytes []byte) {
	bytes, _ = json.Marshal(xRespBody)
	return
}
