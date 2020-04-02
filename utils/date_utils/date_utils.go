package date_utils

import "time"

const (
	dateFormat = "2006-02-01T15:00:55Z"
	apiDBLayout = "2006-02-01"
)

func GetNowDate() time.Time {
	return time.Now().UTC()
}

func GetDateString() string {
	return GetNowDate().Format(dateFormat)
}

func GetNowDBFormat() string {
	return GetNowDate().Format(apiDBLayout)
}
