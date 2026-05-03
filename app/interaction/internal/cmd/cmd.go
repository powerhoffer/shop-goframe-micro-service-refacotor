package cmd

import (
	"context"
	"interaction/internal/controller/collection_info"
	"interaction/internal/controller/comment_info"
	"interaction/internal/controller/praise_info"

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
			collection_info.Register(s)
			comment_info.Register(s)
			praise_info.Register(s)
			s.Run()
			return nil
		},
	}
)
