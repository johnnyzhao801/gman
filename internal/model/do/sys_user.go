// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta     `orm:"table:sys_user, do:true"`
	ID         interface{} //
	Password   interface{} //
	Name       interface{} // 用户姓名
	UserName   interface{} // 用户账户
	Phone      interface{} // 手机号
	Email      interface{} // 邮箱
	CreateTime *gtime.Time //
	WriteTime  *gtime.Time //
	DeleteTime *gtime.Time //
}
