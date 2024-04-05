package util

import "time"

const (
	LayoutDateDDMMYTYY string = "02/01/2006"
)

func ConvertStringToTime(dateStr, layout string) (time.Time, error) {
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
