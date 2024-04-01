package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/shiqinfeng1/goMono/app/gomono-gateway/cmd"
	_ "go.uber.org/automaxprocs"
)

func main() {
	ctx := gctx.GetInitCtx()
	// 如果有子命令，在下面的AddCommand中注册
	err := cmd.Main.AddCommand()
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(ctx)
}
