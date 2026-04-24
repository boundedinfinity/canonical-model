package physical_product

// https://en.wikipedia.org/wiki/Display_resolution

type ScreenDimensions struct {
	Id     ider.Id `json:"id,omitempty"`
	Width  int     `json:"legal-name,omitempty"`
	Height int     `json:"trade-name,omitempty"`
}
