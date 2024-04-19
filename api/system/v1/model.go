package v1

import "github.com/gogf/gf/v2/os/gtime"

type RoleInfo struct {
	Name       string `json:"name"`
	Status     int `json:"status"`
	Info       string `json:"info"`
	UserCount  int `json:"user_count"`
	CreateTime *gtime.Time `json:"create_time"`
	WriteTime  *gtime.Time `json:"write_time"`
}
