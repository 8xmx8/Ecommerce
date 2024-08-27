package main

import (
	_ "Ecommerce/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"Ecommerce/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
