// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	ID         int         `json:"iD"         ` //
	Password   string      `json:"password"   ` //
	Name       string      `json:"name"       ` // 用户姓名
	UserName   string      `json:"userName"   ` // 用户账户
	Phone      string      `json:"phone"      ` // 手机号
	Email      string      `json:"email"      ` // 邮箱
	CreateTime *gtime.Time `json:"createTime" ` //
	WriteTime  *gtime.Time `json:"writeTime"  ` //
	DeleteTime *gtime.Time `json:"deleteTime" ` //
}
