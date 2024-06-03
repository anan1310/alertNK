package request

type ReqUser struct {
	PageInfo
	UserId   int    `json:"userId" `
	Username string `json:"userName" ` // 用户登录名
}
