package time

// Month represents a month of the year, such as "January", "February", etc. It can be represented as an interface to allow for different implementations (e.g., using an enum, a string, or a custom struct).
type Month struct {
	Name   string         `json:"name,omitempty"`
	System CalendarSystem `json:"system,omitempty"`
}
