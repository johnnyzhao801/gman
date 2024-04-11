package service

import (
	"context"
	systemV1 "gman/api/system/v1"
	"gman/internal/model"

	"github.com/casbin/casbin/v2"
)

type (
	ISystem interface {
		GetEnforcer(ctx context.Context) (en *casbin.SyncedEnforcer)
		// SignIn 登录功能
		SignIn(ctx context.Context, in model.UserSignInInput) (err error)
		RoleSearch(ctx context.Context) []*systemV1.RoleInfo
	}
)

var (
	localSystem ISystem
)

func System() ISystem {
	if localSystem == nil {
		panic("implement not found for interface ISystem, forgot register?")
	}
	return localSystem
}

func RegisterSystem(i ISystem) {
	localSystem = i
}
