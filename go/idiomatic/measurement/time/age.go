package time

// Age represents time ages like "iron age", "bronze age", "stone age", etc. It includes a name and a date range.
type Age struct {
	Name  string    `json:"name,omitempty"`
	Range DateRange `json:"range,omitempty"`
}
