package command

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"github.com/shiqinfeng1/goMono/app/biz-trainer/internal/domain/hour"
	"github.com/shiqinfeng1/goMono/app/common/decorator"
)

// 该结构定义‘取消训练’的命令参数
type CancelTraining struct {
	Hour time.Time
}

// 定义命令的操作接口
type CancelTrainingHandler decorator.CommandHandler[CancelTraining]

// 定义实现命令接口的结构体实例
type cancelTrainingHandler struct {
	hourRepo hour.CmdRepo // 使用repo接口，不关心adapters如何实现该接口
}

func NewCancelTrainingHandler(
	hourRepo hour.CmdRepo,
	logger log.Logger,
) CancelTrainingHandler {
	if hourRepo == nil {
		panic("nil hourRepo")
	}
	// ApplyCommandDecorators 对命令以装饰器模式进行了包装，加上了测量和日志的功能
	return decorator.ApplyCommandDecorators[CancelTraining](
		cancelTrainingHandler{hourRepo: hourRepo},
		logger,
	)
}

// 实现 CancelTrainingHandler 接口
func (h cancelTrainingHandler) Handle(ctx context.Context, cmd CancelTraining) error {
	if err := h.hourRepo.UpdateHour(ctx, cmd.Hour, func(hr *hour.Hour) (*hour.Hour, error) {
		// hr是从数据库读出数据厚对应的领域层实例，使用领域层的方法对数据进行更新，更新的数据由UpdateHour负责更新到数据库
		if err := hr.CancelTraining(); err != nil {
			return nil, err
		}
		return hr, nil
	}); err != nil {
		return v1.ErrorCancelTrainingFail("cancel training failed").WithCause(err)
	}
	return nil
}
