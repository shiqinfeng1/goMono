package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/shiqinfeng1/goMono/app/biz-trainer/cmd"
	_ "go.uber.org/automaxprocs"
)

func main() {
	ctx := gctx.GetInitCtx()
	err := cmd.Main.AddCommand(cmd.Service, cmd.JobTrainer, cmd.CronTrainer)
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(ctx)
}
