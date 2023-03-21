package main

import (
	_ "gf-demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
