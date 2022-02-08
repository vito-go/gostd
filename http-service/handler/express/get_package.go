package express

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vito-go/gostd/http-service/handler"
	"github.com/vito-go/gostd/pkg/resp"
)

type GetPackage struct{}

func NewGetPackage() *GetPackage {
	return &GetPackage{}
}

type getPackageData struct {
	BoxId int  `json:"box_id"`
	Exist bool `json:"exist"`
}

func (h *GetPackage) Handle(ctx *gin.Context) {
	handler.Handle(ctx, h)
}

func (h *GetPackage) GetParam(ctx *gin.Context) (*handler.ReqParam, error) {
	return handler.GetParamByPostForm(ctx, "express_id", "package_number")
}

var mockPacakgeNumberMap = map[string]int{"11112222": 1, "33334444": 2, "55556666": 5}

func (h *GetPackage) GetRespBody(ctx *gin.Context, req *handler.ReqParam) *resp.HTTPBody {
	packageNumber := req.Get("package_number")
	boxID, ok := mockPacakgeNumberMap[packageNumber]
	return resp.DataOK(ctx, getPackageData{
		BoxId: boxID,
		Exist: ok,
	})
}

func (h *GetPackage) WriteRespBody(ctx *gin.Context, respBody *resp.HTTPBody) {
	ctx.JSON(http.StatusOK, respBody)
}
