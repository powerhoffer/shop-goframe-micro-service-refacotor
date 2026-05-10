package cmd

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/gateway-resource/internal/controller/file"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway-admin server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/backend", func(group *ghttp.RouterGroup) {
					group.Bind(
						file.NewV1().UploadImage,
					)
				})
				// 需要JWT验证的路由
				group.Group("/backend", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTAuth)
					group.Bind()
				})
			})

			s.Run()
			return nil
		},
	}
)
