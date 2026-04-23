package paper_book

// https://en.wikipedia.org/wiki/ISBN

type InternationalStandardBookNumber struct {
	Group      int `json:"group"`
	Publisher  int `json:"publisher"`
	TItle      int `json:"title"`
	CheckDigit int `json:"check-digit"`
}
