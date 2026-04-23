package iso216

type PaperSize struct {
	Name   string
	Width  int
	Height int
}

type paperSizes struct {
	All []PaperSize
}

var PaperSizes = paperSizes{
	All: []PaperSize{
		// https://papersizes.io/a/a0
		{
			Name:   "A0",
			Width:  841,
			Height: 1189,
		},
		// https://papersizes.io/a/a1
		{
			Name:   "A1",
			Width:  594,
			Height: 841,
		},
		// https://papersizes.io/a/a2
		{
			Name:   "A2",
			Width:  420,
			Height: 594,
		},
		// https://papersizes.io/a/a3
		{
			Name:   "A2",
			Width:  297,
			Height: 420,
		},
	},
}
