package system

import (
	"context"
	"gman/api/system/v1"
	"gman/internal/model"
	"gman/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	err = service.System().SignIn(ctx, model.UserSignInInput{
		UserName: req.UserName,
		Password: req.PassWord,
	})
	return
}
