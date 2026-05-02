package time

type DateRange struct {
	Start Date `json:"start"`
	End   Date `json:"end"`
}

func (this DateRange) IsZero() bool {
	return this.Start.IsZero() && this.End.IsZero()
}

func (this DateRange) In(target Date) bool {
	return !this.Before(target) && !this.After(target)
}

func (this DateRange) Before(target Date) bool {
	if !this.Start.IsZero() && this.Start.Before(target) {
		return true
	}

	return false
}

func (this DateRange) After(target Date) bool {
	if !this.End.IsZero() && this.End.After(target) {
		return true
	}

	return false
}
