package training

import (
	"context"
	"fmt"

	"github.com/shiqinfeng1/goMono/app/biz-training/internal/application/query"
)

// 公共错误，所有实现Repository接口的实例，都会用到该错误
type NotFoundError struct {
	TrainingUUID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("training '%s' not found", e.TrainingUUID)
}

// 定义repo接口，在领域层中定义实现依赖倒置
type Repository interface {
	AddTraining(ctx context.Context, tr *Training) error

	GetTraining(ctx context.Context, trainingUUID string, user User) (*Training, error)

	UpdateTraining(
		ctx context.Context,
		trainingUUID string,
		user User,
		updateFn func(ctx context.Context, tr *Training) (*Training, error),
	) error

	AllTraining(ctx context.Context) ([]query.Training, error)
	FindTrainingForUser(ctx context.Context, userUUID string) ([]query.Training, error)
}
