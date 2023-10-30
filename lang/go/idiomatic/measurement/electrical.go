package measurement

type Electrical struct {
	Capacity string `json:"capacity,omitempty"`
	Voltage  string `json:"voltage,omitempty"`
	Amp      string `json:"amp,omitempty"`
}
