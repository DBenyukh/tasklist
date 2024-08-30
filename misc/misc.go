package misc

import (
	"errors"
	"time"
)

// ToRFC3339 преобразует time.Time в строку формата RFC3339
func ToRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseRFC3339 парсит строку в формате RFC3339 в time.Time
func ParseRFC3339(value string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, errors.New("invalid date format, expected RFC3339")
	}
	return parsedTime, nil
}
