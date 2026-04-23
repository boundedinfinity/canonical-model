package time

import "github.com/boundedinfinity/rfc3339date"

type DateRange struct {
	Start rfc3339date.Rfc3339Date `json:"start"`
	End   rfc3339date.Rfc3339Date `json:"end"`
}
