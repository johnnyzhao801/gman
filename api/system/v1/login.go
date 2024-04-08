package v1

import "github.com/gogf/gf/v2/frame/g"

// LoginReq 登录
type LoginReq struct {
	g.Meta   `path:"/login" tags:"登入登出功能" method:"post" summary:"登录"`
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}
type LoginRes struct {
	g.Meta `mime:"application/json"`
}

// LogoutReq 登出
type LogoutReq struct {
	g.Meta `path:"/logout" method:"get" tags:"登入登出功能" summary:"登出"`
}
type LogoutRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"`
}
