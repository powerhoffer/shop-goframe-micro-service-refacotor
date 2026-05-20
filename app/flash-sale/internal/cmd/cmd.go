package cmd

import (
	"context"

	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/controller/flash_sale"
	logic "shop-goframe-micro-service-refacotor/app/flash-sale/internal/logic/flash_sale"
	"shop-goframe-micro-service-refacotor/app/flash-sale/utility/rabbitmq"
	flashRedis "shop-goframe-micro-service-refacotor/app/flash-sale/utility/redis"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var Main = gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "flash-sale grpc service",
	Func: func(ctx context.Context, parser *gcmd.Parser) error {
		if err := flashRedis.Init(ctx); err != nil {
			g.Log().Errorf(ctx, "秒杀Redis初始化失败: %v", err)
			return err
		}
		if err := rabbitmq.Init(ctx); err != nil {
			g.Log().Warningf(ctx, "秒杀RabbitMQ初始化失败，将跳过消息发布: %v", err)
		}
		if err := logic.PreheatStock(ctx); err != nil {
			g.Log().Errorf(ctx, "秒杀库存预热失败: %v", err)
			return err
		}

		c := grpcx.Server.NewConfig()
		c.Options = append(c.Options, grpcx.Server.ChainUnary(
			grpcx.Server.UnaryValidate,
		))
		s := grpcx.Server.New(c)
		flash_sale.Register(s)
		s.Run()
		return nil
	},
}
