package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/vito-go/logging/tid"
	"github.com/vito-go/mylog"

	"github.com/vito-go/gostd/pkg/resp"
)

// queryType 查询的类型.
type queryType string

const (
	query    queryType = "query"
	postForm queryType = "postForm"
	param    queryType = "param"
)

// Handler 实现Controller接口即可添加路由
// 除了实现接口的四个函数，顺序写在每一个具体的controller文件中。 、
// 其他功能函数，例如组装结果composeUserInfo等一律不可导出使用小写。
type Handler interface {
	Handle(ctx *gin.Context) // MetricHandle 提供添加路由的接口函数
	// GetParam
	// ReqMethod() server.Method                                       // ReqMethod 该 Handler 请求的方法，GET， POST ...
	GetParam(ctx *gin.Context) (*ReqParam, error)               // GetParam 校验以及获取参数
	GetRespBody(ctx *gin.Context, req *ReqParam) *resp.HTTPBody // GetRespBody 获取需要响应的httpBody. 重点聚焦在这个函数的实现
	WriteRespBody(ctx *gin.Context, respBody *resp.HTTPBody)    // WriteRespBody 向前端返回响应内容. 一般是接口返回的http code都是200
}

// ReqParam 请求的参数接口.
type ReqParam struct {
	queryMap map[string]string // queryMap 存储从前端获取到的一些参数
	header   http.Header
	body     []byte
	// 这里可以添加其他项目可能需要的字段，尽管加 向前兼容
	// Auth Auth
}

func (r *ReqParam) QueryMap() map[string]string {
	if r == nil {
		return nil
	}
	return r.queryMap
}
func (r *ReqParam) Body() []byte {
	if r == nil {
		return nil
	}
	return r.body
}

func (r *ReqParam) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Get 获取参数值.
func (r *ReqParam) Get(key string) string {
	return r.queryMap[key]
}

// Set 设定参数值.
func (r *ReqParam) Set(key string, value string) {
	if r.queryMap == nil {
		r.queryMap = make(map[string]string)
	}
	r.queryMap[key] = value
}

// Handle 提供添加路由的接口函数.一个完整的路由请求函数。
// 响应时间、哨兵监控
func Handle(ctx *gin.Context, h Handler) {
	ctx.Set("tid", tid.Get())
	startTime := time.Now()
	reqParam, err := h.GetParam(ctx)
	var httpBody *resp.HTTPBody
	if err != nil {
		httpBody = resp.Err(ctx, err.Error())
	} else {
		httpBody = h.GetRespBody(ctx, reqParam)
	}
	ctx.Writer.Header().Set("access-control-allow-origin", "*")
	h.WriteRespBody(ctx, httpBody)
	var q interface{}
	// 优先从body中取参数
	if reqBody := reqParam.Body(); len(reqBody) != 0 {
		q = reqBody
	} else {
		q = reqParam.QueryMap()
	}
	defer func() {
		mylog.Ctx(ctx).WithField("RT", time.Since(startTime).String()).WithFields(
			"remote_addr", ctx.Request.RemoteAddr,
			"method", ctx.Request.Method,
			"path", ctx.Request.URL.Path,
			"header", reqParam.Header(),
			"query", q,
			"respBody", httpBody).Info("♥")
	}()
}

// GetParamByQuery 通过 ctx.Query 方法获取参数
func GetParamByQuery(ctx *gin.Context, keys ...string) (*ReqParam, error) {
	return checkAndGetParam(ctx, query, keys...)
}

// DefaultParamMap 非必须的参数 key为必传参数名称，value为默认参数值
type DefaultParamMap = map[string]string

// GetParamByQueryWithDefaultParam 通过 ctx.Query 方法获取参数 defaultParamMap 传入非必须的参数 keys为必传参数
func GetParamByQueryWithDefaultParam(ctx *gin.Context, d DefaultParamMap, keys ...string) (*ReqParam, error) {
	return getParamWithDefaultParam(ctx, query, d, keys...)
}

// GetParamByPostFormWithDefaultParam 通过 ctx.PostForm 方法获取参数 defaultParamMap 传入非必须的参数 keys为必传参数
func GetParamByPostFormWithDefaultParam(ctx *gin.Context, d DefaultParamMap, keys ...string) (*ReqParam, error) {
	return getParamWithDefaultParam(ctx, postForm, d, keys...)
}

// GetParamByPostForm 通过 ctx.Request.PostForm方法获取参数
func GetParamByPostForm(ctx *gin.Context, keys ...string) (*ReqParam, error) {
	return checkAndGetParam(ctx, postForm, keys...)
}

// GetParamByParam 通过 ctx.Param 方法获取参数
func GetParamByParam(ctx *gin.Context, keys ...string) (*ReqParam, error) {
	return checkAndGetParam(ctx, param, keys...)
}

func getParamWithDefaultParam(ctx *gin.Context, q queryType, defaultParamMap map[string]string, keys ...string) (*ReqParam, error) {
	reqParam, err := checkAndGetParam(ctx, q, keys...)
	if err != nil {
		return nil, err
	}
	for k, v := range defaultParamMap {
		reqParam.Set(k, v)
		// 如果参数存在按照获取到的参数，如果不存在，设定默认值
		var qv string
		switch q {
		case query:
			qv = ctx.Query(k)
		case postForm:
			qv = ctx.PostForm(k)
		}
		if qv != "" {
			reqParam.Set(k, qv)
		}
	}
	return reqParam, nil
}
func checkAndGetParam(ctx *gin.Context, q queryType, keys ...string) (*ReqParam, error) {
	var hd http.Header
	if err := ctx.BindHeader(&hd); err != nil {
		return nil, err
	}

	var queryMap = make(map[string]string)
	for _, k := range keys {
		var v string
		switch q {
		case query:
			v = ctx.Query(k)
		case postForm:
			v = ctx.PostForm(k)
		case param:
			v = ctx.Param(k)
		default:
			return nil, errors.New("unknown queryType")
		}
		if v == "" {
			return nil, fmt.Errorf("%s 参数错误", k)
		}
		queryMap[k] = v
	}

	reqParam := &ReqParam{
		queryMap: queryMap,
		header:   hd,
	}
	return reqParam, nil
}
