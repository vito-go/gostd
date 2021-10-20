package userinfo

type Interface interface {
	GetUserInfoByID(id int64) *UserInfo
	Hello() string
}
