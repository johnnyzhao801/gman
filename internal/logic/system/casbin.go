package system

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	casbinadapter "github.com/hailaz/gf-casbin-adapter/v2"
	"sync"
)

var (
	once sync.Once
	myDB gdb.DB
	en   *casbin.SyncedEnforcer
)

func (s *sSystem) GetEnforcer(ctx context.Context) *casbin.SyncedEnforcer {
	once.Do(func() {
		en = initEnforcer(ctx)
	})
	return en
}

func initEnforcer(ctx context.Context) (en *casbin.SyncedEnforcer) {
	var err error
	link := g.Cfg().MustGet(ctx, "database.default.link")
	rbacConf := g.Cfg().MustGet(ctx, "casbin.modelFile").String()
	myDB, err = gdb.New(gdb.ConfigNode{
		Link: link.String(),
	})
	if err != nil {
		panic(err)
	}
	a := casbinadapter.NewAdapter(casbinadapter.Options{GDB: myDB})
	en, err = casbin.NewSyncedEnforcer(rbacConf, a)
	if err != nil {
		panic(err)
	}
	err = en.LoadPolicy()
	if err != nil {
		panic(err)
	}
	return
}
