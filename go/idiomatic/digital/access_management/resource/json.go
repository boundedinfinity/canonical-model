package resource

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrResourceJson = errorer.New("JSON")
	errResourceJson = errorer.Func(ErrResource, ErrResourceJson)
)

type unmarshalResource struct {
	Kind Kind `json:"kind"`
}

func UnmarshalJSON(data []byte) (Resource, error) {
	var resource Resource
	var err error
	var unmarshaled unmarshalResource

	if err = json.Unmarshal(data, &unmarshaled); err != nil {
		return resource, err
	}

	switch unmarshaled.Kind {
	case Kinds.Person:
		var person Person

		if err = json.Unmarshal(data, &person); err != nil {
			err = errResourceJson("unmarshal error : %w", err)
		} else {
			resource = &person
		}
	default:
		err = errResourceJson("unmarshal error : unsupported resource: %s", unmarshaled.Kind)
	}

	return resource, err
}

type marshalResource struct {
	Kind Kind `json:"kind"`
	Data any  `json:"data"`
}

func MarshalJSON(resource Resource) ([]byte, error) {
	var data []byte
	var err error

	marshaled := marshalResource{
		Kind: resource.GetKind(),
	}

	switch d := resource.GetKind(); d {
	case Kinds.Person:
		marshaled.Data = d
	default:
		err = errResourceJson("marshal error : unsupported resource: %s", resource.GetKind())
	}

	if err != nil {
		data, err = json.Marshal(marshaled)
	}

	return data, err
}
