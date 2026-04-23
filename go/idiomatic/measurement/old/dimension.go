package measurement

type Dimensions[T ~int] struct {
	Width    Length[T] `json:"width,omitempty"`
	Heigth   Length[T] `json:"heigth,omitempty"`
	Depth    Length[T] `json:"depth,omitempty"`
	Length   Length[T] `json:"length,omitempty"`
	Radius   Length[T] `json:"radius,omitempty"`
	Diameter Length[T] `json:"diameter,omitempty"`
}
