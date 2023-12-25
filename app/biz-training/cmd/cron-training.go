package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	CronTraining = &gcmd.Command{
		Name:        "cron-training",
		Brief:       "start training cron task",
		Description: "entry to start a training cron task",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// todo: start cron here
			return nil
		},
	}
	//todo: add more cron-xxx here
)
