package v1

import "github.com/gogf/gf/v2/os/gtime"

type RoleInfo struct {
	Name       string
	Status     int
	Info       string
	UserCount  int
	CreateTime *gtime.Time
	WriteTime  *gtime.Time
}
