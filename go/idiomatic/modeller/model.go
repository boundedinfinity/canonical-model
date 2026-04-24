package modeller

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

type Kind interface {
	~string
}

type Modeller interface {
	GetId() ider.Id
	GetName() string
}
