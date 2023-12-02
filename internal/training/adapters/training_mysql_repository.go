package adapters

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/shiqinfeng1/goMono/internal/common/config"
	"github.com/shiqinfeng1/goMono/internal/training/application/query"
	"github.com/shiqinfeng1/goMono/internal/training/domain/training"

	"github.com/go-kratos/kratos/v2/log"
)

type trainingRepo struct {
	adpr *config.Adapter
	log  *log.Helper
	db   *sqlx.DB
}

// NewTrainingRepo .
func NewTrainingRepo(adpr *config.Adapter, logger log.Logger) training.Repository {
	db, err := sqlx.Connect(adpr.Database.Driver, adpr.Database.Source)
	if err != nil {
		return nil
	}
	return &trainingRepo{
		adpr: adpr,
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

func (r trainingRepo) AllTraining(ctx context.Context) ([]query.Training, error) {
	return nil, nil
}

func (r trainingRepo) FindTrainingForUser(ctx context.Context, userUUID string) ([]query.Training, error) {
	return nil, nil
}

// warning: RemoveAllTraining was designed for tests for doing data cleanups
func (r trainingRepo) RemoveAllTraining(ctx context.Context) error {
	return nil
}
