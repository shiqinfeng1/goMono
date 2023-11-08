package hour

import "github.com/pkg/errors"

// 定义小时的状态
var (
	Available         = Availability{"available"}          // 可用：未来的未被安排的hour
	NotAvailable      = Availability{"not_available"}      // 不可用：过去的hour
	TrainingScheduled = Availability{"training_scheduled"} // 已安排训练：过去或未来已被安排训练计划的
)

var availabilityValues = []Availability{
	Available,
	NotAvailable,
	TrainingScheduled,
}

// Availability is enum.
//
// Using struct instead of `type Availability string` for enums allows us to ensure,
// that we have full control of what values are possible.
// 如果你使用 `type Availability string` 定义，那么y就可以把任何未定义的字符串转为Availability： `Availability("i_can_put_anything_here")`
type Availability struct {
	a string
}

// 类型转换：字符串转为Availability
func NewAvailabilityFromString(availabilityStr string) (Availability, error) {
	for _, availability := range availabilityValues {
		if availability.String() == availabilityStr {
			return availability, nil
		}
	}
	return Availability{}, errors.Errorf("unknown '%s' availability", availabilityStr)
}

// 检查是否为空
func (h Availability) IsZero() bool {
	return h == Availability{}
}

// 类型转换：Availability转为字符串
func (h Availability) String() string {
	return h.a
}

var (
	ErrTrainingScheduled   = errors.New("unable to modify hour, because scheduled training")
	ErrNoTrainingScheduled = errors.New("training is not scheduled")
	ErrHourNotAvailable    = errors.New("hour is not available")
)

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
