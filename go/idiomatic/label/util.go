package label

import (
	"strings"

	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func getName(model LabelModel) string {
	return model.Name
}

func getId(model LabelModel) ider.Id {
	return model.Id
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func equalsAny(target LabelModel) func(LabelModel) bool {
	id := equalsId(target)
	name := equalsName(target)

	return func(model LabelModel) bool {
		return id(model) || name(model)
	}
}

func equalsName(target LabelModel) func(LabelModel) bool {
	return func(model LabelModel) bool {
		return strings.EqualFold(target.Name, model.Name)
	}
}

func equalsId(target LabelModel) func(LabelModel) bool {
	return func(model LabelModel) bool {
		return target.Id == model.Id &&
			!target.Id.IsZero() &&
			!model.Id.IsZero()
	}
}
