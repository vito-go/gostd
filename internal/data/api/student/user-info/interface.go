package userinfo

type Interface interface {
	GetUserInfoMapByID(id int64) (map[string]string, error)
	GetNameByID(id int64) (string, error)
	Hello() string
}
