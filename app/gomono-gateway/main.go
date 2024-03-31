package main

import (
	_ "net/http/pprof"

	_ "github.com/go-kratos/gateway/middleware/bbr"
	_ "github.com/go-kratos/gateway/middleware/cors"
	_ "github.com/go-kratos/gateway/middleware/logging"
	_ "github.com/go-kratos/gateway/middleware/rewrite"
	_ "github.com/go-kratos/gateway/middleware/tracing"
	_ "github.com/go-kratos/gateway/middleware/transcoder"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/cmd"
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
