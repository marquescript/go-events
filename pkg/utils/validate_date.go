package utils

import "time"

func ParseDate(date string) (time.Time, error) {
	dateParsed, err := time.Parse("2006-01-02", date)
	return dateParsed, err
}
