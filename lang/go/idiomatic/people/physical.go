package people

type Physical struct {
	Height   float32 `json:"height,omitempty"`
	Weight   float32 `json:"weight,omitempty"`
	EyeColor string  `json:"eye-color,omitempty"`
	Sex      string  `json:"sex,omitempty"`
}
