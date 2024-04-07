package user

import (
	"context"
	"sims/internal/service"
	"sims/api/user/v1"
)

func (c *ControllerV1) UserSearch(ctx context.Context, req *v1.UserSearchReq) (res *v1.UserSearchRes, err error) {
	res = &v1.UserSearchRes{
		UserList: service.User().UserSearch(ctx),
	}
	return
}
