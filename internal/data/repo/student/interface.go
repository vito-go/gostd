package student

import (
	"context"

	helloblogdao "github.com/vito-go/gostd/internal/data/dao/helloblog-dao"
)

// UseInfoAPI 所有提供给logic曾的api接口方法在此定义。函数名，参数、返回值清晰。
// logic层开发人员通过此interface来沟通交流数据
// 此为后续构建自动化api文档做下铺垫，从此告别手动更新维护api文档.
// 开发顺序： 先定义接口，再来实现
type UseInfoAPI interface {
	// GetUserInfoByProfile 根据profile获取所有的信息 返回结果字段后期调整。
	GetUserInfoByProfile(ctx context.Context, profile string) (*helloblogdao.UserInfoModel, error)
}
