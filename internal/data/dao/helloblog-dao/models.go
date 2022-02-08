package helloblogdao

import (
	"time"
)

// UserInfoModel user_info表.
type UserInfoModel struct {
	ID           int64         `json:"id,omitempty"`
	Phone        int64         `json:"phone,omitempty"`
	Password     string        `json:"password,omitempty"`
	Nick         string        `json:"nick,omitempty"`
	Sex          string        `json:"sex,omitempty"` //  用户性别， 男，女，未知
	Email        string        `json:"email,omitempty"`
	Avatar       string        `json:"avatar,omitempty"`
	Birthday     time.Time     `json:"birthday"`
	Hometown     string        `json:"hometown,omitempty"`
	Company      string        `json:"company,omitempty"`
	Job          string        `json:"job,omitempty"`
	Technology   []string      `json:"technology,omitempty"`
	Website      string        `json:"website,omitempty"`
	WechatQrcode string        `json:"wechat_qrcode,omitempty"`
	Intro        string        `json:"intro,omitempty"`
	Slogan       string        `json:"slogan,omitempty"`
	RegisterTime time.Time     `json:"register_time"`
	Profile      string        `json:"profile,omitempty"`
	Theme        userInfoTheme `json:"theme,omitempty"`
	BlogScore    int64         `json:"blog_score,omitempty"`
	Weight       float32       `json:"weight,omitempty"`
}

// comment on column user_info.phone  is '用户手机号';
// comment on column user_info.password  is '加盐密码';
// comment on column user_info.password  is '用户昵称';
// comment on column user_info.sex  is '用户性别， 男，女，未知';
// comment on column user_info.email  is '用户邮箱';
// comment on column user_info.avatar  is '用户头像(链接)';
// comment on column user_info.birthday  is '用户生日';
// comment on column user_info.hometown  is '家乡（例如 河南郑州）';
// comment on column user_info.company  is '公司';
// comment on column user_info.job  is '职业';
// comment on column user_info.technology  is '技术栈 例如 {"go","python"]';
// comment on column user_info.website  is '用户个人网站';
// comment on column user_info.wechat_qrcode  is '用户个人微信二维码';
// comment on column user_info.intro  is '用户个人简介';
// comment on column user_info.slogan  is '用户个性签名，也叫口号，用户个人主页中用到';
// comment on column user_info.register_time  is '注册时间';
// comment on column user_info.profile  is '个人主页 id';
// comment on column user_info.theme  is '主题';
// comment on column user_info.weight  is '个人权重';分';

type userInfoTheme int
