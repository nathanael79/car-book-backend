package utils

import (
	"time"
)

func ConvertStringToDate(dateString string) (time.Time, error) {
	const layout = "2006-01-02"
	parsedDate, err := time.Parse(layout, dateString)

	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}
