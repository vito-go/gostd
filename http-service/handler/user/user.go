// Package user
// 联系人接口文档.https://xxx.com
package user

// User 用户相关接口
type User struct {
	// ////////
	// studentRepo *student.Repo
	GetUserInfo *GetUserInfo
}

func NewUser(GetUserInfo *GetUserInfo) *User {
	return &User{GetUserInfo: GetUserInfo}
}
