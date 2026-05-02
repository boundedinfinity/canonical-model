package time

type Date struct{}

func (this Date) IsZero() bool {
	return true
}

func (this Date) Before(date Date) bool {
	return true
}

func (this Date) After(date Date) bool {
	return true
}

func (this Date) Equal(date Date) bool {
	return true
}
