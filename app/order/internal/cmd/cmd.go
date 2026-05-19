package cmd

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/order/internal/controller/order_info"
	"shop-goframe-micro-service-refacotor/app/order/internal/controller/refund_info"
	"shop-goframe-micro-service-refacotor/app/order/utility/payment"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "order grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			if err := payment.InitWechatClient(); err != nil {
				g.Log().Errorf(ctx, "支付客户端初始化失败:%v", err)
				return err
			}

			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			order_info.Register(s)
			refund_info.Register(s)
			s.Run()
			return nil
		},
	}
)
