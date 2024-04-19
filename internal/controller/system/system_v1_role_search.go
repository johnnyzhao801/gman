package system

import (
	"context"

	"gman/api/system/v1"
	"gman/internal/service"
)

func (c *ControllerV1) RoleSearch(ctx context.Context, req *v1.RoleSearchReq) (res *v1.RoleSearchRes, err error) {
	res = &v1.RoleSearchRes{
		RoleList: service.System().RoleSearch(ctx),
	}
	return 
}
