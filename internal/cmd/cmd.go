package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
	"gman/internal/controller/system"
	"gman/internal/controller/user"
	"gman/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetSessionStorage(gsession.NewStorageRedisHashTable(g.Redis()))
			s.BindMiddlewareDefault(ghttp.MiddlewareCORS)
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Group("/system", func(group *ghttp.RouterGroup) {
					group.Middleware(ghttp.MiddlewareHandlerResponse)
					group.Bind(
						system.NewV1(),
					)
				})
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Middleware(ghttp.MiddlewareHandlerResponse)
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						user.NewV1(),
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
