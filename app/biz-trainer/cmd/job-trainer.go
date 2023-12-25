package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	JobTrainer = &gcmd.Command{
		Name:        "job-trainer",
		Brief:       "start trainer job",
		Description: "entry to start a trainer job",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// todo: start job here
			return nil
		},
	}
	//todo: add more job-xxx here
)
