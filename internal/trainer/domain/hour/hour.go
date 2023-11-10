package hour

import (
	"time"
)

// 训练课程以小时为单位，这里将小时作为一个领域实体，Hour具有一个具体的时间，还有一个状态
// 领域实体的数据只存储于内存中
type Hour struct {
	hour time.Time

	availability Availability
}

func (h *Hour) Time() time.Time {
	return h.hour
}

func (h Hour) Availability() Availability {
	return h.availability
}

func (h Hour) IsAvailable() bool {
	return h.availability == Available
}

func (h Hour) HasTrainingScheduled() bool {
	return h.availability == TrainingScheduled
}

// 已经被安排训练的Hour是不能被设置为不可用的
func (h *Hour) MakeNotAvailable() error {
	if h.HasTrainingScheduled() {
		return ErrTrainingScheduled
	}

	h.availability = NotAvailable
	return nil
}

// 已经被安排训练的Hour是不能被设置为可用的
func (h *Hour) MakeAvailable() error {
	if h.HasTrainingScheduled() {
		return ErrTrainingScheduled
	}

	h.availability = Available
	return nil
}

// 不可用的Hour是不能被安排训练
func (h *Hour) ScheduleTraining() error {
	if !h.IsAvailable() {
		return ErrHourNotAvailable
	}

	h.availability = TrainingScheduled
	return nil
}

// 只有已被安排训练的Hour才能取消训练
func (h *Hour) CancelTraining() error {
	if !h.HasTrainingScheduled() {
		return ErrNoTrainingScheduled
	}

	h.availability = Available
	return nil
}
