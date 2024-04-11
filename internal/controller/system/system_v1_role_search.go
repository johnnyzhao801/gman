package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gman/api/system/v1"
)

func (c *ControllerV1) RoleSearch(ctx context.Context, req *v1.RoleSearchReq) (res *v1.RoleSearchRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
