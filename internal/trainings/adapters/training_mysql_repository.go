package adapters

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/shiqinfeng1/goMono/internal/trainings/app/query"
	"github.com/shiqinfeng1/goMono/internal/trainings/conf"
	"github.com/shiqinfeng1/goMono/internal/trainings/domain/training"

	"github.com/go-kratos/kratos/v2/log"
)

type trainingRepo struct {
	data *conf.Adapter
	log  *log.Helper
	db   *sqlx.DB
}

// NewTrainingRepo .
func NewTrainingRepo(data *conf.Adapter, logger log.Logger) training.Repository {
	db, err := sqlx.Connect(data.Database.Driver, data.Database.Source)
	if err != nil {
		return nil
	}
	return &trainingRepo{
		data: data,
		log:  log.NewHelper(logger),
		db:   db,
	}
}

func (r trainingRepo) AddTraining(ctx context.Context, tr *training.Training) error {
	return nil
}

func (r trainingRepo) GetTraining(
	ctx context.Context,
	trainingUUID string,
	user training.User,
) (*training.Training, error) {
	return nil, nil
}

func (r trainingRepo) UpdateTraining(
	ctx context.Context,
	trainingUUID string,
	user training.User,
	updateFn func(ctx context.Context, tr *training.Training) (*training.Training, error),
) error {
	return nil
}

func (r trainingRepo) AllTrainings(ctx context.Context) ([]query.Training, error) {
	return nil, nil
}

func (r trainingRepo) FindTrainingsForUser(ctx context.Context, userUUID string) ([]query.Training, error) {
	return nil, nil
}

// warning: RemoveAllTrainings was designed for tests for doing data cleanups
func (r trainingRepo) RemoveAllTrainings(ctx context.Context) error {
	return nil
}
