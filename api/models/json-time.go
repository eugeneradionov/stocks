package models

import (
	"strings"
	"time"
)

const JSONTimeFormat = time.RFC3339

type JSONTime time.Time

func (t JSONTime) Time() time.Time {
	return time.Time(t)
}

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}

	str := t.Time().Format(JSONTimeFormat)
	return []byte(str), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	parsed, err := time.Parse(JSONTimeFormat, s)
	if err != nil {
		return err
	}

	*t = JSONTime(parsed)

	return nil
}
