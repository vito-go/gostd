package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vito-go/mylog"

	"github.com/vito-go/gostd/api"
	"github.com/vito-go/gostd/http-service/handler"
	"github.com/vito-go/gostd/internal/conn"
	"github.com/vito-go/gostd/pkg/resp"
)

type GetUserInfo struct {
	UserCli *api.UserCli
}

func NewGetUserInfo(userRpcCli *conn.UserRpcCli) *GetUserInfo {
	return &GetUserInfo{UserCli: api.NewUserCli(userRpcCli)}
}

func (h *GetUserInfo) Handle(ctx *gin.Context) {
	handler.Handle(ctx, h)
}

func (h *GetUserInfo) GetParam(ctx *gin.Context) (*handler.ReqParam, error) {
	return handler.GetParamByParam(ctx, "profile")
}

func (h *GetUserInfo) GetRespBody(ctx *gin.Context, req *handler.ReqParam) *resp.HTTPBody {
	profile := req.Get("profile")
	result, err := h.UserCli.GetUserInfoByProfile(profile)
	if err != nil {
		mylog.Ctx(ctx).WithField("profile", profile).Error(err.Error())
		// return resp.Err(ctx, httperr.ErrInternal.Error()) todo 应该不对端上直接暴露err
		return resp.Err(ctx, err.Error())
	}
	mylog.Ctx(ctx).WithFields("profile", profile, "result", result).Info("UserCli.GetUserInfoByProfile -->")
	return resp.DataOK(ctx, result)
}

func (h *GetUserInfo) WriteRespBody(ctx *gin.Context, respBody *resp.HTTPBody) {
	ctx.JSON(http.StatusOK, respBody)
}
