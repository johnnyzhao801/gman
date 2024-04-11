// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Name       string      `json:"name"       ` // 角色名
	Status     int         `json:"status"     ` // 状态
	UserCount  int         `json:"userCount"  ` // 用户数量
	Info       string      `json:"info"       ` // 角色信息
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	WriteTime  *gtime.Time `json:"writeTime"  ` // 更新时间
}
