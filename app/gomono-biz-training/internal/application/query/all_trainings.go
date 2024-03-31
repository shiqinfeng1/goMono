package query

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/utils/decorator"
)

type AllTraining struct{}

type AllTrainingHandler decorator.QueryHandler[AllTraining, []Training]

type allTrainingHandler struct {
	readModel AllTrainingReadModel
}

func NewAllTrainingHandler(
	readModel AllTrainingReadModel,
	logger log.Logger,
) AllTrainingHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[AllTraining, []Training](
		allTrainingHandler{readModel: readModel},
		logger,
	)
}

type AllTrainingReadModel interface {
	AllTraining(ctx context.Context) ([]Training, error)
}

func (h allTrainingHandler) Handle(ctx context.Context, _ AllTraining) (tr []Training, err error) {
	return h.readModel.AllTraining(ctx)
}
