package resp

import (
	"context"
	"fmt"
)

const (
	generalErrCode = 1000000
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrParse = Error("系统内部数据解析错误")
)

type HTTPBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Tid     int64       `json:"tid,omitempty"` // Tid traceID 全链路追踪tid
}

// DataOK data should, can be marshaled.
func DataOK(ctx context.Context, data interface{}) *HTTPBody {
	tid, _ := ctx.Value("tid").(int64)
	return &HTTPBody{
		Data: data,
		Tid:  tid,
	}
}
func Err(ctx context.Context, message string) *HTTPBody {
	tid, _ := ctx.Value("tid").(int64)
	return &HTTPBody{
		Code:    generalErrCode,
		Message: message,
		Tid:     tid,
	}
}
func Errf(ctx context.Context, format string, args ...interface{}) *HTTPBody {
	tid, _ := ctx.Value("tid").(int64)
	return &HTTPBody{
		Code:    generalErrCode,
		Message: fmt.Sprintf(format, args...),
		Tid:     tid,
	}
}
