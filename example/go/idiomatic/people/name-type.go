package people

import "github.com/boundedinfinity/schema/idiomatic/id"

type NameType struct {
	Id          id.Id  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
