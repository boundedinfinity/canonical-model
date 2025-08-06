package people

// https://www.tailorstore.com/guide/how-to/how-to-measure-guide

type Measurements struct {
	Height        float32 `json:"height,omitempty"`
	Weight        float32 `json:"weight,omitempty"`
	EyeColor      string  `json:"eye-color,omitempty"`
	Sex           string  `json:"sex,omitempty"`
	RingSize      string  `json:"ring-size,omitempty"`
	Waist         string  `json:"waist,omitempty"`
	Neck          string  `json:"neck,omitempty"`
	Inseam        string  `json:"inseam,omitempty"`
	Chest         string  `json:"chest,omitempty"`
	Cup           string  `json:"cup,omitempty"`
	Hip           string  `json:"hip,omitempty"`
	Wrist         string  `json:"wrist,omitempty"`
	ShirtLength   string  `json:"shirt-length,omitempty"`
	ArmLength     string  `json:"arm-length,omitempty"`
	ShoulderWidth string  `json:"shoulder-width,omitempty"`
}
