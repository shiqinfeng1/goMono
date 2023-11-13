package query

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/common/auth"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
)

type TrainingForUser struct {
	User auth.User
}

type TrainingForUserHandler decorator.QueryHandler[TrainingForUser, []Training]

type trainingForUserHandler struct {
	readModel TrainingForUserReadModel
}

func NewTrainingForUserHandler(
	readModel TrainingForUserReadModel,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) TrainingForUserHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[TrainingForUser, []Training](
		trainingForUserHandler{readModel: readModel},
		logger,
		metricsClient,
	)
}

type TrainingForUserReadModel interface {
	FindTrainingForUser(ctx context.Context, userUUID string) ([]Training, error)
}

func (h trainingForUserHandler) Handle(ctx context.Context, query TrainingForUser) (tr []Training, err error) {
	return h.readModel.FindTrainingForUser(ctx, query.User.UUID)
}
