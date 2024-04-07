package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "sims/internal/packed"

	_ "sims/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"sims/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
