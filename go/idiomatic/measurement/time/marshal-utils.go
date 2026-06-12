package time

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

func MarshalJSON(s fmt.Stringer) ([]byte, error) {
	return json.Marshal(s.String())
}

func UnmarshalJSON[V ~int, I ~string](v *V, data []byte, parser func(I) (V, error)) error {
	var i I
	var err error

	if err = json.Unmarshal(data, &i); err != nil {
		return err
	}

	*v, err = parser(i)

	return err
}

func Value(v fmt.Stringer) (driver.Value, error) {
	return v.String(), nil
}

func Scan[V ~int, I ~string | ~int](v *V, value any, parser func(I) (V, error)) error {
	var dv driver.Value
	var i I
	var ok bool
	var err error

	switch value.(type) {
	case string:
		dv, err = driver.String.ConvertValue(value)
	case int:
		dv, err = driver.Int32.ConvertValue(value)
	default:
		return fmt.Errorf("cannot convert %T to string", value)
	}

	if err != nil {
		return err
	}

	if i, ok = dv.(I); !ok {
		return errInvalidTimeFormatFn("not a string")
	}

	*v, err = parser(i)
	return err
}
