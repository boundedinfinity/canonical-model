package taxes

type Tax struct {
	Year        int     `json:"year"`
	Tax         float64 `json:"tax"`
	Asseessment float64 `json:"assessment"`
}
