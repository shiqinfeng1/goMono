package hour

import (
	"time"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

// Hour工厂的配置，可以配置生成Hour领域实体的一些策略和参数限制
type FactoryConfig struct {
	MaxWeeksInTheFutureToSet int
	MinUtcHour               int
	MaxUtcHour               int
}

// 检查Hour工厂的配置的有效性
// 使用multierr包， 检查全部参数后，再一起返回可能的全部错误
func (f FactoryConfig) Validate() error {
	var err error

	if f.MaxWeeksInTheFutureToSet < 1 {
		err = multierr.Append(
			err,
			errors.Errorf(
				"MaxWeeksInTheFutureToSet should be greater than 1, but is %d",
				f.MaxWeeksInTheFutureToSet,
			),
		)
	}
	if f.MinUtcHour < 0 || f.MinUtcHour > 24 {
		err = multierr.Append(
			err,
			errors.Errorf(
				"MinUtcHour should be value between 0 and 24, but is %d",
				f.MinUtcHour,
			),
		)
	}
	if f.MaxUtcHour < 0 || f.MaxUtcHour > 24 {
		err = multierr.Append(
			err,
			errors.Errorf(
				"MinUtcHour should be value between 0 and 24, but is %d",
				f.MaxUtcHour,
			),
		)
	}

	if f.MinUtcHour > f.MaxUtcHour {
		err = multierr.Append(
			err,
			errors.Errorf(
				"MinUtcHour (%d) can't be after MaxUtcHour (%d)",
				f.MinUtcHour, f.MaxUtcHour,
			),
		)
	}

	return err
}

// Hour工厂
type Factory struct {
	// it's better to keep FactoryConfig as a private attribute,
	// thanks to that we are always sure that our configuration is not changed in the not allowed way
	fc FactoryConfig
}

// 以指定的config生成一个Hour工厂
func NewFactory(fc FactoryConfig) (Factory, error) {
	if err := fc.Validate(); err != nil {
		return Factory{}, errors.Wrap(err, "invalid config passed to factory")
	}

	return Factory{fc: fc}, nil
}

// 给测试用例使用，不允许构造失败
func MustNewFactory(fc FactoryConfig) Factory {
	f, err := NewFactory(fc)
	if err != nil {
		panic(err)
	}

	return f
}

func (f Factory) Config() FactoryConfig {
	return f.fc
}

func (f Factory) IsZero() bool {
	return f == Factory{}
}

// 接口只需要传通用的参数，避免对外暴露自定义的类型，因此这里封装2个行为函数，NewAvailableHour和NewNotAvailableHour， 而不是把availability作为参数传递
// Hour工厂生产一个可用的Hour，
func (f Factory) NewAvailableHour(hour time.Time) (*Hour, error) {
	if err := f.validateTime(hour); err != nil {
		return nil, err
	}

	return &Hour{
		hour:         hour,
		availability: Available,
	}, nil
}

// Hour工厂生产一个不可用的Hour
func (f Factory) NewNotAvailableHour(hour time.Time) (*Hour, error) {
	if err := f.validateTime(hour); err != nil {
		return nil, err
	}

	return &Hour{
		hour:         hour,
		availability: NotAvailable,
	}, nil
}

// UnmarshalHourFromDatabase unmarshals Hour from the database.
//
// It should be used only for unmarshalling from the database!
// You can't use UnmarshalHourFromDatabase as constructor - It may put domain into the invalid state!
// 把数据库中的数据（字段）转换为领域实体结构，并且只能用于转化数据库的数据转换， 不能用于构造新的Hour
func (f Factory) UnmarshalHourFromDatabase(hour time.Time, availability Availability) (*Hour, error) {
	// 由工厂进行参数校验
	if err := f.validateTime(hour); err != nil {
		return nil, err
	}

	if availability.IsZero() {
		return nil, errors.New("empty availability")
	}

	return &Hour{
		hour:         hour,
		availability: availability,
	}, nil
}

func (f Factory) validateTime(hour time.Time) error {
	if !hour.Round(time.Hour).Equal(hour) { //把hour四舍五入精确到小时，例如一小时40分会返回2小时
		return ErrNotFullHour
	}

	// AddDate is better than Add for adding days, because not every day have 24h!
	if hour.After(time.Now().AddDate(0, 0, f.fc.MaxWeeksInTheFutureToSet*7)) {
		return TooDistantDateError{
			MaxWeeksInTheFutureToSet: f.fc.MaxWeeksInTheFutureToSet,
			ProvidedDate:             hour,
		}
	}

	currentHour := time.Now().Truncate(time.Hour) // 把当前时间值精确到小时，例如一小时40分会返回1小时
	if hour.Before(currentHour) || hour.Equal(currentHour) {
		return ErrPastHour
	}
	if hour.UTC().Hour() > f.fc.MaxUtcHour {
		return TooLateHourError{
			MaxUtcHour:   f.fc.MaxUtcHour,
			ProvidedTime: hour,
		}
	}
	if hour.UTC().Hour() < f.fc.MinUtcHour {
		return TooEarlyHourError{
			MinUtcHour:   f.fc.MinUtcHour,
			ProvidedTime: hour,
		}
	}

	return nil
}
