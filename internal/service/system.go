// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sims/internal/model"

	"github.com/casbin/casbin/v2"
)

type (
	ISystem interface {
		GetEnforcer(ctx context.Context) (en *casbin.SyncedEnforcer)
		// SignIn creates session for given userhahah account.
		SignIn(ctx context.Context, in model.UserSignInInput) (err error)
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
