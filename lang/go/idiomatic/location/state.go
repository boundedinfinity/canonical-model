package location

type State struct {
	Name StateName `json:"name,omitempty"`
	Code StateAnsi `json:"code,omitempty"`
}

func (t State) String() string {
	return t.Name.String()
}
