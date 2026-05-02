package label

import (
	"strings"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func getName(model Label) string {
	return model.Name
}

func getId(model Label) ider.Id {
	return model.Id
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func equalsAny(target Label) func(Label) bool {
	id := equalsId(target)
	name := equalsName(target)

	return func(model Label) bool {
		return id(model) || name(model)
	}
}

func equalsName(target Label) func(Label) bool {
	return func(model Label) bool {
		return strings.EqualFold(target.Name, model.Name)
	}
}

func equalsId(target Label) func(Label) bool {
	return func(model Label) bool {
		return target.Id == model.Id &&
			!target.Id.IsZero() &&
			!model.Id.IsZero()
	}
}
