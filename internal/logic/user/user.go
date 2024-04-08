package user

import (
	"context"
	"fmt"
	"gman/internal/dao"
	"gman/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	userV1 "gman/api/user/v1"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) IsSignedIn(ctx context.Context) bool {
	r := g.RequestFromCtx(ctx)
	v := r.Session.MustGet("account")
	fmt.Println(v)
	return !v.IsNil()
}

func (s *sUser) UserSearch(ctx context.Context) []*userV1.UserInfo {
	var res []*userV1.UserInfo
	userModel := dao.SysUser.Ctx(ctx)
	_ = userModel.FieldsEx(dao.SysUser.Columns().Password).Scan(&res)
	return res
}
