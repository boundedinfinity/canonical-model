package time

import (
	"database/sql/driver"

	"github.com/boundedinfinity/go-commoner/errorer"
)

type TimeFormat int

var (
	ErrInvalidTimeFormat   = errorer.New("invalid time format")
	errInvalidTimeFormatFn = errorer.Func(ErrInvalidTimeFormat)
)

func ParseTimeFormat(s string) (TimeFormat, error) {
	switch s {
	case "measurement.time.time-format.iso8601":
		return TimeFormats.Iso8601, nil
	case "measurement.time.time-format.rfc3339":
		return TimeFormats.Rfc3339, nil
	case "measurement.time.time-format.unix":
		return TimeFormats.Unix, nil
	case "measurement.time.time-format.numeric":
		return TimeFormats.Numeric, nil
	case "measurement.time.time-format.canonical":
		return TimeFormats.Canonical, nil
	default:
		return 0, errInvalidTimeFormatFn(s)
	}
}

func (this TimeFormat) String() string {
	switch this {
	case TimeFormats.Iso8601:
		return "measurement.time.time-format.iso8601"
	case TimeFormats.Rfc3339:
		return "measurement.time.time-format.rfc3339"
	case TimeFormats.Unix:
		return "measurement.time.time-format.unix"
	case TimeFormats.Numeric:
		return "measurement.time.time-format.numeric"
	case TimeFormats.Canonical:
		return "measurement.time.time-format.canonical"
	default:
		return errInvalidTimeFormatFn("%d", this).Error()
	}
}

func (this TimeFormat) IsZero() bool {
	return this == 0
}

func (this TimeFormat) MarshalJSON() ([]byte, error) {
	return MarshalJSON(this)
}

func (this *TimeFormat) UnmarshalJSON(data []byte) error {
	return UnmarshalJSON(this, data, ParseTimeFormat)
}

func (this TimeFormat) Value() (driver.Value, error) {
	return Value(this)
}

func (this *TimeFormat) Scan(value any) error {
	return Scan(this, value, ParseTimeFormat)
}

var TimeFormats = timeFormats{
	Iso8601:   1,
	Rfc3339:   2,
	Unix:      3,
	Numeric:   4,
	Canonical: 5,
}

type timeFormats struct {
	Iso8601   TimeFormat
	Rfc3339   TimeFormat
	Unix      TimeFormat
	Numeric   TimeFormat
	Canonical TimeFormat
}
