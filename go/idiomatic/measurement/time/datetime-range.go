package time

import "github.com/boundedinfinity/rfc3339date"

type DateTimeRange struct {
	Start rfc3339date.Rfc3339DateTime `json:"start"`
	End   rfc3339date.Rfc3339DateTime `json:"end"`
}
