package time

type TimeRange struct {
	Start Time `json:"start"`
	End   Time `json:"end"`
}

func (this TimeRange) IsZero() bool {
	return this.Start.IsZero() && this.End.IsZero()
}

func (this TimeRange) In(target Time) bool {
	return !this.Before(target) && !this.After(target)
}

func (this TimeRange) Before(target Time) bool {
	if !this.Start.IsZero() && this.Start.Before(target) {
		return true
	}

	return false
}

func (this TimeRange) After(target Time) bool {
	if !this.End.IsZero() && this.End.After(target) {
		return true
	}

	return false
}
