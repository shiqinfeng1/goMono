package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/cmd"
	_ "go.uber.org/automaxprocs"
)

func main() {
	ctx := gctx.GetInitCtx()
	err := cmd.Main.AddCommand()
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(ctx)
}
