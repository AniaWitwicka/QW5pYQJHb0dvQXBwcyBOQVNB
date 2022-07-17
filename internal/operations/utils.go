package operations

import (
	"fmt"
	"time"
)

func ParseDate(s string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ParseAndValidateDateRanges(from, to string) (time.Time, time.Time, error) {
	fromDate, err := ParseDate(from)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("Error while parsing from date: %s", err)
	}
	toDate, err := ParseDate(to)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("Error while parsing to date: %s", err)
	}
	if toDate.Before(fromDate) {
		return time.Time{}, time.Time{}, fmt.Errorf("Wrong date ranges")
	}
	return fromDate, toDate, nil
}

func GetUrlLen(from, to time.Time) int {
	return int(to.Sub(from).Hours()/24) + 1
}
