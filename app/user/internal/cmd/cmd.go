package cmd

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/user/internal/controller/consignee_info"
	"shop-goframe-micro-service-refacotor/app/user/internal/controller/user_info"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "user grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			consignee_info.Register(s)
			user_info.Register(s)
			s.Run()
			return nil
		},
	}
)
