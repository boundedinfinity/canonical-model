package label

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrLabelUnmarshalJSON = errors.New("invalid label unmarshal")

type labelDescriminator struct {
	Type  string          `json:"type,omitempty"`
	Label json.RawMessage `json:"label,omitempty"`
}

type labelWrapper struct {
	Type  string     `json:"type,omitempty"`
	Label ValueLabel `json:"label,omitempty"`
}

func LabelMarshalJSON(v ValueLabel) ([]byte, error) {
	wrapper := labelWrapper{
		Label: v,
		Type:  v.Type(),
	}

	return json.Marshal(wrapper)
}

func LabelUnmarshalJSON(data []byte) (ValueLabel, error) {
	var err error
	var descriminator labelDescriminator
	var label ValueLabel

	if err = json.Unmarshal(data, &descriminator); err != nil {
		return nil, err
	}

	switch descriminator.Type {
	case NameOnlyLabel{}.Type():
		label = &NameOnlyLabel{}
	case StringLabel{}.Type():
		label = &StringLabel{}
	case IntegerLabel{}.Type():
		label = &IntegerLabel{}
	case FloatLabel{}.Type():
		label = &FloatLabel{}
	case DateLabel{}.Type():
		label = &DateLabel{}
	case DateTimeLabel{}.Type():
		label = &DateTimeLabel{}
	default:
		err = fmt.Errorf("%w for %v", ErrLabelUnmarshalJSON, string(data))
		return label, err
	}

	err = json.Unmarshal(descriminator.Label, &label)
	return label, err
}
