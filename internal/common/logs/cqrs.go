package logs

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
)

func LogCommandExecution(commandName string, cmd interface{}, err error) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	l := log.NewHelper(logger)
	if err == nil {
		l.Info(commandName + " command succeeded")
	} else {
		l.Error(commandName + " command failed")
	}
}
