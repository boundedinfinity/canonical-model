package time

import "github.com/boundedinfinity/rfc3339date"

type TimeRange struct {
	Start rfc3339date.Rfc3339Time `json:"start"`
	End   rfc3339date.Rfc3339Time `json:"end"`
}
