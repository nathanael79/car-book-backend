package utils

import "time"

func ConvertStringToDateTime(dateString string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", dateString)

	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
