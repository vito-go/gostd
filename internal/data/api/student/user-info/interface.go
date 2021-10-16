package userinfo

type Interface interface {
	GetUserInfoById(id int64) *UserInfo
	Hello() string
}
