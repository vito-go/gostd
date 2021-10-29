package userinfo

// Interface 所有提供给logic曾的api接口方法在此定义。函数名，参数、返回值清晰。
// logic层开发人员通过此interface来沟通交流数据
// 此为后续构建自动化api文档做下铺垫，从此告别手动更新维护api文档.
type Interface interface {
	GetUserInfoMapByID(id int64) (map[string]string, error)
	GetNameByID(id int64) (string, error)
	Hello() string
}
