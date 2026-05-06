package event

type Event interface {
	Kind() Kind
	String() string
}

// type Event struct {
// 	Id                  ider.Id               `json:"id,omitempty,omitzero"`
// 	Kind                Kind                  `json:"kind,omitempty"`
// 	Title               string                `json:"title"`
// 	Description         string                `json:"description"`
// 	Notes               []note.Note           `json:"notes"`
// 	Mandatory           []person.Person       `json:"mandatory"`
// 	Optional            []person.Person       `json:"optional"`
// 	Awareness           []person.Person       `json:"awareness"`
// 	DateRange           time.DateRange        `json:"date-range"`
// 	TimeRange           time.TimeRange        `json:"time-range"`
// 	Labels              label.Labels          `json:"labels"`
// 	Visibility          visibility.Visibility `json:"visibility"`
// 	PreparationDuration time.Duration         `json:"preparation-duration"`
// 	PostDuration        time.Duration         `json:"post-duration"`
// 	Location            location.Location     `json:"location"`
// 	Recurring           time.Interval         `json:"recurring,omitempty"`
// }
