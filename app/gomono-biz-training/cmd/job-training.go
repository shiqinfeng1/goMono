package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	JobTraining = &gcmd.Command{
		Name:        "job-training",
		Brief:       "start training job",
		Description: "entry to start a training job",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// todo: start job here
			return nil
		},
	}
	//todo: add more job-xxx here
)
