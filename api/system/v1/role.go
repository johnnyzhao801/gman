package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RoleSearchReq struct {
	g.Meta `path:"/list" method:"get" tags:"角色管理" summary:"系统角色列表"`
}

type RoleSearchRes struct {
	g.Meta   `mime:"application/json"`
	RoleList []*RoleInfo `json:"role_list"`
}
