// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"gman/api/user/v1"
)

type IUserV1 interface {
	UserSearch(ctx context.Context, req *v1.UserSearchReq) (res *v1.UserSearchRes, err error)
}
