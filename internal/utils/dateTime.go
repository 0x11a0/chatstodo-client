package utils

import (
	"log"
	"time"
)

const (
	GOOGLE_CALENDAR_FORMAT = "2006-01-02T15:04:05+08:00"
	BACKEND_FORMAT         = "2006-01-02T15:04:05.000Z"
	DATETIME_HTML_FORMAT   = "2006-01-02T15:04"
	// DD/MM/YYYY - HH:MM
	DATETIME_PRETTY_FORMAT = "02/01/2006 - 15:04"
)

var TIMEZONES = map[string]string{
	"Singapore": "Asia/Singapore",
}

func HTMLToTime(dateTime string) *time.Time {
	resultTime, err := time.Parse(DATETIME_HTML_FORMAT, dateTime)
	if err != nil {
		log.Println("dateTime.go - HTMLToGCalendar()")
		log.Println(err)
		return nil
	}
	return &resultTime
}

func HTMLToGCalendar(dateTime string) string {
	resultTime, err := time.Parse(DATETIME_HTML_FORMAT, dateTime)
	if err != nil {
		log.Println("dateTime.go - HTMLToGCalendar()")
		log.Println(err)
		return ""
	}
	return resultTime.Format(GOOGLE_CALENDAR_FORMAT)
}

// Parses a date in ISO8601 format.
// Returns the Time object if valid or nil if
// an error occurs
func ParseISOString(dateTime string) *time.Time {
	resultTime, err := time.Parse(BACKEND_FORMAT, dateTime)
	if err != nil {
		log.Println("dateTime.go - ParseISO()")
		log.Println("PROBLEM DATE", dateTime)
		log.Println(err)
		return nil
	}
	return &resultTime
}

// Returns the string converted into the specified
// local timezone, in the format specified for
// html input datetime-local tag. If countryName is not found,
// defaults to Singapore. Returns empty string
// if error occurs or if input dateTime is nil.
func GetLocalDateTimeDatePicker(dateTime *time.Time,
	countryName string) string {
	if dateTime == nil {
		return ""
	}
	timezone := TIMEZONES[countryName]
	if timezone == "" {
		// Defaults to Singapore timezone
		timezone = TIMEZONES["Singapore"]
	}

	localLocation, err := time.LoadLocation(timezone)
	if err != nil {
		log.Println("dateTime.go - GetLocalDateTimeDatePicker(), load location")
		return ""
	}

	return dateTime.In(localLocation).Format(DATETIME_HTML_FORMAT)
}

// Returns the string converted into the specified
// local timezone, in the format "DD/MM/YYYY - HH:MM" for
// pretty displaying. If countryName is not found,
// defaults to Singapore. Returns empty string
// if error occurs or if input dateTime is nil.
func GetLocalDateTimePretty(dateTime *time.Time,
	countryName string) string {
	if dateTime == nil {
		return ""
	}
	timezone := TIMEZONES[countryName]
	if timezone == "" {
		// Defaults to Singapore timezone
		timezone = TIMEZONES["Singapore"]
	}

	localLocation, err := time.LoadLocation(timezone)
	if err != nil {
		log.Println("dateTime.go - GetLocalDateTimePretty(), load location")
		return ""
	}

	return dateTime.In(localLocation).Format(DATETIME_PRETTY_FORMAT)
}

// Converts datetime string in html datetime-local
// format to "DD/MM/YYYY - HH:MM" for pretty display.
// Returns empty string if error occurs while parsing
func PrettifyHTMLDateTime(htmlDateTime string) string {
	dateTime, err := time.Parse(DATETIME_HTML_FORMAT, htmlDateTime)

	if err != nil {
		log.Println("dateTime.go - PrettifyHTMLDateTime(), parse time")
		return ""
	}
	return dateTime.Format(DATETIME_PRETTY_FORMAT)
}
