package utils

import (
	"fmt"
	"time"
)

// GetYesterdayDate get the date of yesterday.
func GetYesterdayDate() string {
	currentTs := time.Now()
	date2Yesterday := currentTs.AddDate(0, 0, -1)
	return fmt.Sprintf("%s", date2Yesterday.Format("2006-01-02"))
}

// GetCurrentDateTime get the current time
func GetCurrentDateTime() string {
	currentTs := time.Now()
	return fmt.Sprintf("%s", currentTs.Format("2006-01-02 15:04:05"))
}
