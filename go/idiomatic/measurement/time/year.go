package time

import "fmt"

type Year int

func (this Year) String() string {
	return fmt.Sprintf("%d", this)
}

func (this Year) IsLeap() bool {
	if this%400 == 0 {
		return true
	}

	if this%100 == 0 {
		return false
	}

	return this%4 == 0
}
