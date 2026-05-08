package time

type Year struct {
	Value  int            `json:"value,omitempty"`
	System CalendarSystem `json:"system,omitempty"`
}
