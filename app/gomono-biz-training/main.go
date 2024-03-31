package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-training/cmd"
	_ "go.uber.org/automaxprocs"
)

func main() {
	ctx := gctx.GetInitCtx()
	err := cmd.Main.AddCommand(cmd.Service, cmd.JobTraining, cmd.CronTraining)
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(ctx)
}
