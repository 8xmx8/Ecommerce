package main

import (
	"Ecommerce/internal/cmd"
	_ "Ecommerce/internal/logic"
	_ "Ecommerce/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
