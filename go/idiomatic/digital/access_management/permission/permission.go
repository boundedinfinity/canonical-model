package permission

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

var _ ider.Idable = &PermissionModel{}

type PermissionModel struct {
	ID          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
}

func (this PermissionModel) GetId() ider.Id {
	return this.ID
}
