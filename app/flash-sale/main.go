package main

import (
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	conf, err := g.Cfg().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}

	grpcx.Resolver.Register(etcd.New(conf.String()))
	cmd.Main.Run(ctx)
}
