package password

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

type RequirementsModel struct {
	Id                ider.Id  `json:"id"`
	Name              string   `json:"name"`
	Size              int      `json:"size"`
	Alhpa             bool     `json:"alpha"`
	AlhpaLower        bool     `json:"alpha-lower"`
	AlhpaUpper        bool     `json:"alpha-upper"`
	Numeric           bool     `json:"numeric"`
	Symbols           bool     `json:"symbols"`
	BeginsWithAlpha   bool     `json:"begins-with-alpha"`
	NoRepeatedSqences bool     `json:"no-repeated-sequences"`
	NotAllowed        []string `json:"not-allowed"`
}
