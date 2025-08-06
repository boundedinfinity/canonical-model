package calver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// https://calver.org/

type CalVer struct {
	Year  int
	Month int
	Day   int
}

func (t CalVer) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	_, err := buf.WriteString(fmt.Sprintf("%4d-%2d-%2d", t.Year, t.Month, t.Day))

	return buf.Bytes(), err
}

func (t *CalVer) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	ymd := stringer.Split(s, "-")

	if len(ymd) != 3 {
		return fmt.Errorf("invalid calendar version: %v", ymd)
	}

	if y, err := strconv.Atoi(ymd[0]); err != nil {
		return err
	} else {
		t.Year = y
	}

	if m, err := strconv.Atoi(ymd[1]); err != nil {
		return err
	} else {
		t.Month = m
	}

	if d, err := strconv.Atoi(ymd[2]); err != nil {
		return err
	} else {
		t.Day = d
	}

	return nil
}
