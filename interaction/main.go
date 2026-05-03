package main

import (
	_ "interaction/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"interaction/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
