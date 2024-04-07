package system

import (
	"context"
	"sims/api/system/v1"
	"sims/internal/model"
	"sims/internal/service"
	//"sims/internal/model"
	//"sims/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	err = service.System().SignIn(ctx, model.UserSignInInput{
		UserName: req.Account,
		Password: req.PassWord,
	})
	return
}
