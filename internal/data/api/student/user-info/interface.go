package userinfo

type Interface interface {
	GetUserInfoMapByID(id int64) (map[string]string, error)
	GetNameById(id int64) (string, error)
	Hello() string
}
