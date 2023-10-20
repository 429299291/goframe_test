package main

import (
	_ "golangt/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"golangt/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
