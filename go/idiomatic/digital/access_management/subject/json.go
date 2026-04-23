package subject

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrSubjectJson = errorer.New("JSON")
	errSubjectJson = errorer.Func(ErrSubject, ErrSubjectJson)
)

type unmarshalSubject struct {
	Kind Kind `json:"kind"`
}

func UnmarshalJSON(data []byte) (Subject, error) {
	var subject Subject
	var err error
	var unmarshaled unmarshalSubject

	if err = json.Unmarshal(data, &unmarshaled); err != nil {
		return subject, err
	}

	switch unmarshaled.Kind {
	case Kinds.Person:
		var person Person

		if err = json.Unmarshal(data, &person); err != nil {
			err = errSubjectJson("unmarshal error : %w", err)
		} else {
			subject = &person
		}
	default:
		err = errSubjectJson("unmarshal error : unsupported subject: %s", unmarshaled.Kind)
	}

	return subject, err
}

type marshalSubject struct {
	Kind Kind `json:"kind"`
	Data any  `json:"data"`
}

func MarshalJSON(subject Subject) ([]byte, error) {
	var data []byte
	var err error

	marshaled := marshalSubject{
		Kind: subject.GetKind(),
	}

	switch d := subject.GetKind(); d {
	case Kinds.Person:
		marshaled.Data = d
	default:
		return data, errSubjectJson("marshal error : unsupported subject: %s", subject.GetKind())
	}

	data, err = json.Marshal(marshaled)

	return data, err
}
