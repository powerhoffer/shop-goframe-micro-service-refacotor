package cmd

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/banner"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/goods"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/interaction"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/order"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/user"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway-h5 server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/frontend", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.NewV1().UserInfoLogin,
						user.NewV1().UserInfoRegister,
					)
				})
				// 需要JWT验证的路由
				group.Group("/frontend", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTAuth)
					group.Bind(
						// 需要认证的接口
						user.NewV1().UserInfo,
						user.NewV1().UserInfoUpdatePassword,
						user.NewV1().ConsigneeInfoCreate,
						user.NewV1().ConsigneeInfoGetList,
						user.NewV1().ConsigneeInfoUpdate,
						user.NewV1().ConsigneeInfoDelete,
						interaction.NewV1(),
						goods.NewV1(),
						order.NewV1(),
						banner.NewV1(),
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
