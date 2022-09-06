package main

import (
	_ "site/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"site/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
