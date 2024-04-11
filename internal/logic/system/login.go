package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"gman/internal/dao"
	"gman/internal/model"
	"gman/internal/model/do"
	"gman/internal/model/entity"
	"gman/internal/service"
)

type (
	sSystem struct{}
)

func init() {
	service.RegisterSystem(New())
}

func New() service.ISystem {
	return &sSystem{}
}

// SignIn 登录功能
func (s *sSystem) SignIn(ctx context.Context, in model.UserSignInInput) (err error) {
	var user *entity.SysUser
	fmt.Println("输入的数据为：", in)
	r := g.RequestFromCtx(ctx)
	ctx = gctx.New()
	// 配置文件中的sessionIdName
	//sessionIdName := a.Map()["sessionIdName"]
	//v.Set(sessionIdName)
	//sessionId := r.Cookie.Get(v.String())
	// 去数据库查询是否存在
	err = dao.SysUser.Ctx(ctx).Where(do.SysUser{
		UserName: in.UserName,
		Password: in.Password,
	}).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "密码错误!")
		return
	}
	if user == nil {
		return gerror.New(`Passport or Password not correct`)
	}
	r.Session.MustSet("account", in)
	//fmt.Println("sessionId为：", sessionId)
	return
}
