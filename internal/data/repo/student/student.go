package student

type Client struct {
	UserInfoRepo UseInfoAPI
}

func NewClient(u *UserInfo) *Client {
	return &Client{UserInfoRepo: u}
}
