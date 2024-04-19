package system

import (
	"context"
	systemV1 "gman/api/system/v1"
	"gman/internal/dao"
)

func (s *sSystem) RoleSearch(ctx context.Context) []*systemV1.RoleInfo {
	var res []*systemV1.RoleInfo
	roleModel := dao.SysRole.Ctx(ctx)
	roleModel.Scan(&res)
	return res
}
