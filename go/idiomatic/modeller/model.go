package modeller

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type Kind interface {
	~string
}

type Modeller interface {
	GetId() ider.Id
	GetName() string
}
