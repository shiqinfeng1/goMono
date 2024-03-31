package types

import (
	"fmt"

	"github.com/pkg/errors"
)

type Mode struct {
	s string
}

var (
	ModeDevelop = Mode{"develop"} //
	ModeTest    = Mode{"test"}    //
	ModeProduct = Mode{"product"} //
)

var modeValues = []Mode{
	ModeDevelop,
	ModeTest,
	ModeProduct,
}

func (m Mode) IsZero() bool {
	return m == Mode{}
}

func (m Mode) String() string {
	return m.s
}

func (m Mode) IsValid() bool {
	for _, v := range modeValues {
		if m.s == v.s {
			return true
		}
	}
	return false
}

func (m Mode) Is(expect Mode) bool {
	return m == expect
}

func (m Mode) ErrInvaild() error {
	return errors.New(fmt.Sprintf("env.mode=%s is invalid, not in %v", m.s, modeValues))
}

func NewModeFromString(modeStr string) Mode {
	return Mode{s: modeStr}
}
