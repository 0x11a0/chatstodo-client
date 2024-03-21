package utils

import (
	"log"
	"time"
)

const (
	ISO_8601_FORMAT       = "2015-01-27T05:57:31.399861+00:00"
	DATETIME_LOCAL_FORMAT = "2024-10-30T17:45"
)

var TIMEZONES = map[string]string{
	"Singapore": "Asia/Singapore",
}

// Parses a date in ISO8601 format.
// Returns the Time object if valid or nil if
// an error occurs
func ParseISOString(dateTime string) *time.Time {
	resultTime, err := time.Parse(ISO_8601_FORMAT, dateTime)
	if err != nil {
		log.Println("dateTime.go - ParseISO()")
		log.Println(err)
		return nil
	}
	return &resultTime
}

// Returns the string converted into the specified
// local timezone, in the format specified for
// html input datetime-local tag. If timeFormat is not found,
// defaults to Singapore. Returns empty string
// if error occurs.
func GetLocalDateTime(dateTime *time.Time,
	countryName string, timeFormat string) string {
	timezone := TIMEZONES[countryName]
	if timezone == "" {
		// Defaults to Singapore timezone
		timezone = TIMEZONES["Singapore"]
	}

	localLocation, err := time.LoadLocation(timezone)
	if err != nil {
		log.Println("calendarApi.go - GetLocalDateTime(), load location")
		return ""
	}

	return dateTime.In(localLocation).Format(DATETIME_LOCAL_FORMAT)
}
