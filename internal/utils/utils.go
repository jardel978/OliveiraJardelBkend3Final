package utils

import "time"

const (
	DATE_FORMAT = "2006-01-02T15:00:00Z"
)

func ParseDate(date string) (dateTime time.Time, err error) {
	dateTime, err = time.Parse(DATE_FORMAT, date)
	if err != nil {
		return dateTime, err
	}
	return dateTime, nil
}

func ParseDataTime(horario string) (dateTime time.Time, err error) {
	t, err := time.Parse(time.RFC3339, horario)
	if err != nil {
		return t, err
	}
	return t, nil
}

func DateIsValid(date string, position string) (bool, error) {
	dateTime, err := time.Parse(DATE_FORMAT, date)
	if err != nil {
		return false, err
	}

	switch position {
	case "after":
		return dateTime.After(time.Now()), nil // data posterior a hoje
	case "before":
		return dateTime.Before(time.Now()), nil // data anterior a hoje
	}
	return false, nil
}
