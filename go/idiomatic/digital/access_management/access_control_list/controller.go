package access_control_list

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/model/digital/access_management/permission"
	"github.com/boundedinfinity/canonical_model/idiomatic/model/digital/access_management/resource"
	"github.com/boundedinfinity/canonical_model/idiomatic/model/util/slicer"
)

func Memory() *Controller {
	return &Controller{
		lists: []AccessControlListModel{},
	}
}

type Controller struct {
	lists []AccessControlListModel
}

func (this Controller) Granted(resource resource.Resource, requested AccessControlLisItem) (bool, error) {
	var hasAllPermissions = true
	var list AccessControlListModel
	var found bool

	for _, list := range this.lists {
		if list.Resource.GetKind() == resource.GetKind() || list.Resource.GetId() == resource.GetId() {
			found = true
			break
		}
	}

	if !found {
		return hasAllPermissions, nil
	}

	for _, item := range list.Items {
		if item.Subject.GetKind() != requested.Subject.GetKind() || item.Subject.GetId() != requested.Subject.GetId() {
			continue
		}

		if !this.permissionsMatch(item.Permissions, requested.Permissions) {
			hasAllPermissions = false
			break
		}
	}

	return hasAllPermissions, nil
}

func (this Controller) permissionsMatch(items []permission.PermissionModel, targets []permission.PermissionModel) bool {
	for _, target := range targets {
		if !slicer.Contains(items, target) {
			return false
		}
	}

	return true
}
