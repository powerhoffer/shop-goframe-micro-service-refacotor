package main

import (
	_ "shop-goframe-micro-service-refacotor/app/search/internal/packed"
	"shop-goframe-micro-service-refacotor/app/search/utility/elasticsearch"

	"github.com/gogf/gf/v2/os/gctx"

	"shop-goframe-micro-service-refacotor/app/search/internal/cmd"
)

func main() {
	ctx := gctx.New()

	// 初始化ES
	if err := elasticsearch.Init(ctx); err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
