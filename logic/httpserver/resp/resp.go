package resp

import (
	"encoding/json"
	"fmt"

	"gitea.com/liushihao/gostd/logic/mylog"
)

const (
	// generalErrCode todo 暂定就一种错误代码
	generalErrCode = 1
)

type respData struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func DataOK(data interface{}) []byte {
	r := respData{
		Data: data,
	}
	return r.toJSONB()
}
func DataErr(message string) []byte {
	r := respData{
		Code:    generalErrCode,
		Message: message,
	}
	return r.toJSONB()
}
func DataErrF(format string, args ...interface{}) []byte {
	r := respData{
		Code:    generalErrCode,
		Message: fmt.Sprintf(format, args...),
	}
	return r.toJSONB()
}
func (r *respData) toJSONB() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		mylog.Error(err)
	}
	return b
}
