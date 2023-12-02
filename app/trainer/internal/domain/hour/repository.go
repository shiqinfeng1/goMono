package hour

import (
	"context"
	"time"
)

// 在领域层定义的repo接口，在adapters层的repo需要实现该接口
type CmdRepo interface {
	// 根据hourTime的时间查询Hour结构，返回的是领域层使用的实体，在repo中负责将数据库对应的结构转换为该实体结构,
	// 转换方式：1，repo自己实现转换（字段较多）；2，domain实现实体生成工厂，由repo调用生产（字段较少）
	GetHour(ctx context.Context, hourTime time.Time) (*Hour, error)
	// 更新数据库的Hour数据，查询的条件是hourTime，更新方法是updateFn，更新后持久化到数据库， 更新操作使用事务。
	UpdateHour(
		ctx context.Context,
		hourTime time.Time,
		updateFn func(h *Hour) (*Hour, error),
	) error
}
