package v1

import "github.com/gogf/gf/v2/frame/g"

// LoginReq 登录
type LoginReq struct {
	g.Meta   `path:"/login" tags:"Hello" method:"post" summary:"You first api"`
	Account  string `json:"account"`
	PassWord string `json:"pass_word"`
}
type LoginRes struct {
	g.Meta `mime:"application/json"`
}

// LogoutReq 登出
type LogoutReq struct {
	g.Meta `path:"/logout" method:"get" tags:"UserService" summary:"Sign out current user"`
}
type LogoutRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"`
}
