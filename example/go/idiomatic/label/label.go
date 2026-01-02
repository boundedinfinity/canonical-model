package label

import (
	"fmt"

	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Label struct {
	Id           id.Id  `json:"id,omitempty,omitzero"`
	Name         string `json:"name,omitempty,omitzero"`
	Abbreviation string `json:"abbreviation,omitempty,omitzero"`
	Description  string `json:"description,omitempty,omitzero"`
}

func (this Label) String() string {
	return fmt.Sprintf("[%s] %s", this.Id, this.Name)
}

func (this Label) IsZero() bool {
	return this.Id.IsZero() &&
		this.Name == "" &&
		this.Abbreviation == "" &&
		this.Description == ""
}
