package main

import (
	_ "shop-goframe-micro-service-refacotor/app/search/internal/packed"
	"shop-goframe-micro-service-refacotor/app/search/utility/binlog"
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
	// 启动时先全量同步一次，保证 ES 中有已有商品数据。
	if err := binlog.SyncAllGoodsToES(ctx); err != nil {
		panic(err)
	}
	// 启动 binlog 监听（在后台运行）
	go binlog.StartBinlogSyncer(ctx)

	cmd.Main.Run(gctx.GetInitCtx())
}
