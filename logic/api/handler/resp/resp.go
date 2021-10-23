package resp

import (
	"encoding/json"
	"fmt"

	"gitea.com/liushihao/mylog"
)

const (
	// todo 暂定就一种错误代码
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
	return r.toJsonB()
}
func DataErr(message string) []byte {
	r := respData{
		Code:    generalErrCode,
		Message: message,
	}
	return r.toJsonB()
}
func DataErrF(format string, args ...interface{}) []byte {
	r := respData{
		Code:    generalErrCode,
		Message: fmt.Sprintf(format, args...),
	}
	return r.toJsonB()
}
func (r *respData) toJsonB() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		mylog.Error(err)
	}
	return b
}
