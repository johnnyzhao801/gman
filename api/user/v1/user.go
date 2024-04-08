package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserSearchReq struct {
	g.Meta `path:"/list" method:"get" tags:"用户管理" summary:"系统用户列表"`
}

type UserSearchRes struct {
	g.Meta   `mime:"application/json"`
	UserList []*UserInfo `json:"user_list"`
}
