package cmd

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/goods_images"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/goods_info"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "interaction grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			goods_images.Register(s)
			goods_info.Register(s)
			s.Run()
			return nil
		},
	}
)
