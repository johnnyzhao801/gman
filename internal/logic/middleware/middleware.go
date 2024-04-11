package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"gman/internal/service"
	"net/http"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	ctx := r.Context()
	e := service.System().GetEnforcer(ctx)
	g.Log().Info(ctx, r.RequestURI)
	g.Log().Info(ctx, r.Method)
	//ok, err := e.DeletePermissionsForUser("admin")
	//g.Log().Debug(ctx, "okle", ok)
	//if err != nil {
	//	return
	//}
	if ok, err := e.Enforce("hailaz", r.RequestURI, "GET"); ok {
		g.Log().Debug(ctx, "okle", ok)
	} else {
		g.Log().Debug(ctx, "errle", err)
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    400,
			Message: "Bad Request",
		})
		return
	}
	if service.User().IsSignedIn(ctx) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
