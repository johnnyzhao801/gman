// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta     `orm:"table:sys_role, do:true"`
	Name       interface{} // 角色名
	Status     interface{} // 状态
	UserCount  interface{} // 用户数量
	Info       interface{} // 角色信息
	CreateTime *gtime.Time // 创建时间
	WriteTime  *gtime.Time // 更新时间
}
