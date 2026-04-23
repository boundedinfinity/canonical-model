package calendar

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/label"
)

type CalendarModel struct {
	Id          ider.Id        `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Labels      []label.Labels `json:"labels"`
}
