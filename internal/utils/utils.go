package utils

import "time"

const (
	DATE_FORMAT = "2006-01-02"
)

func ParseDate(date string) (dateTime time.Time, err error) {
	dateTime, err = time.Parse(DATE_FORMAT, date)
	if err != nil {
		return dateTime, err
	}
	return dateTime, nil
}

func ParseDataTime(horario string) (dateTime time.Time, err error) {
	t, err := time.Parse("2022-12-14T19:07:21.311-03:00", horario)
	if err != nil {
		return t, err
	}
	return t, nil
}
