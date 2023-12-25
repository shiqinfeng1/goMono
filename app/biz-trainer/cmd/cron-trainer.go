package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	CronTrainer = &gcmd.Command{
		Name:        "cron-trainer",
		Brief:       "start trainer cron task",
		Description: "entry to start a trainer cron task",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// todo: start cron here
			return nil
		},
	}
	//todo: add more cron-xxx here
)
