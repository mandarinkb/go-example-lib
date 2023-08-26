package datetime

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Date and Time now
var DateTimeNow = time.Now()

const (
	DefaultTimeZone                       = "Asia/Bangkok"                         // DefaultTimeZone Default Time Zone Thailand
	DefaultDateTimeFormate                = "2006-01-02 15:04:05"                  // set default DateTime format
	DefaultDateTimeFormateFullAsiaBangkok = "2006-01-02 15:04:05.354619 +0700 +07" // set default DateTime format
	DefaultDateTimeFormate0700            = "2006-01-02T15:04:05+07:00"            // set default DateTime format
	DefaultDateTimeFormate070007          = "2006-01-02 15:04:05 +0700 +07"        // set default DateTime format 2006-01-02 15:04:05 +0700 +07 : Asia/Bangkok
	YYYYmmddHHmmSSFormatDatetime          = "20060102150405"                       // YYYYmmddhisFormatDatetime DataTime format YearMonthDayHourMinutesSecond
	YYYYMMDD                              = "2006-01-02"                           // YYYY-MM-DD: 2022-03-23
	HHMMSS24h                             = "15:04:05"                             // 24h hh:mm:ss: 14:23:20
	HHMMSS12h                             = "3:04:05 PM"                           // 12h hh:mm:ss: 2:23:20 PM
	TextDate                              = "January 2, 2006"                      // text date: March 23, 2022
	TextDateWithWeekday                   = "Monday, January 2, 2006"              // text date with weekday: Wednesday, March 23, 2022
	AbbrTextDate                          = "Jan 2 Mon"                            // abbreviated text date: Mar 23 Wed
)

// GetDateTimeCountryZone :: function to get date time from country zone format 2018-01-01 00:00:00
func GetDateTimeCountryZone(strDateTime string) time.Time {
	t, _ := time.Parse(DefaultDateTimeFormate, strDateTime)
	year, month, day := t.Date()
	hour, min, second := t.Clock()
	nsec := t.Nanosecond()
	loc, _ := time.LoadLocation(DefaultTimeZone)

	//set timezone,
	retTime := time.Date(year, month, day, hour, min, second, nsec, loc)
	return retTime
}

// GetCurrentLocationTimeZoneAsiaBangkok :: Current location time zone "AsiaBangkok"
func GetCurrentLocationTimeZoneAsiaBangkok() *time.Location {
	locAt, _ := time.LoadLocation(DefaultTimeZone)
	return locAt
}

// GetCurrentDateTimeNano :: Time nano
func GetCurrentDateTimeNano() string {
	return GetCurrentDateTimeNowToString("2006-01-02 15:04:05.000000000")
}

// GetCurrentDateTimeNow :: GetCurrentTime get current DateTime
func GetCurrentDateTimeNowNotFormat() time.Time {
	locAt, _ := time.LoadLocation(DefaultTimeZone)
	current := time.Now().In(locAt)
	return current
}

// GetCurrentDateTimeNowToString :: GetCurrentTime get current DateTime
func GetCurrentDateTimeNowToString(format string) string {
	if strings.TrimSpace(format) == "" {
		format = DefaultDateTimeFormate
	}
	locAt, _ := time.LoadLocation(DefaultTimeZone)
	current := time.Now().In(locAt).Format(format)
	return current
}

// Get Current Date Time Now return value is ( 2006-01-02 15:04:05.354619 +0700 +07 )
func GetCurrentDateTimeNow() time.Time {
	locAt, _ := time.LoadLocation(DefaultTimeZone)
	strCurrentDateTimeNow := time.Now().UTC().In(locAt)
	return strCurrentDateTimeNow
}

// GetTimeZone :: Time Zone
func GetTimeZone() string {
	loc, _ := time.LoadLocation(DefaultTimeZone)
	localDateTime, err := time.ParseInLocation(DefaultDateTimeFormate, DateTimeNow.Format(DefaultDateTimeFormate), loc)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return localDateTime.Format(DefaultDateTimeFormate)
}

// ParseDateTimeDefaultFormateToString t is parse date time format string, return value is ( 2006-01-02 15:04:05 )
func ParseDateTimeDefaultFormateToString(issuedDate time.Time) string {
	locAt, _ := time.LoadLocation(DefaultTimeZone)
	return issuedDate.In(locAt).Format(DefaultDateTimeFormate)
}

// ParseDateTimeDefaultFormate is parse date time format string, return value is ( 2006-01-02 15:04:05 )
func ParseDateTimeDefaultFormate(issuedDate string) string {
	temp, _ := time.Parse(DefaultDateTimeFormate, issuedDate)
	return temp.Format(DefaultDateTimeFormate)
}

// ParseDateTimeDefaultFormateToDateTime is parse date time format string, return value is ( 2006-01-02 15:04:05 )
func ParseDateTimeByFormateToDateTime(format string, issuedDate time.Time) time.Time {
	if strings.TrimSpace(format) == "" {
		format = DefaultDateTimeFormate
	}
	temp, _ := time.Parse(format, fmt.Sprintf("%v", issuedDate))
	return temp
}

// ParseDateTimeByFormateToDateTimeToString is parse date time format string, return value is ( 2006-01-02 )
func ParseDateTimeByFormateToDateTimeToString(format string, issuedDate time.Time) string {
	if strings.TrimSpace(format) == "" {
		format = DefaultDateTimeFormate
	}
	locAt, _ := time.LoadLocation(DefaultTimeZone)
	return issuedDate.In(locAt).Format(format)
}

// ParseDateTime_CustomFmt is parse date time format string, return value is ( 2006-01-02 15:04:05 )
func ParseDateTimeByFormate(format, issuedDate string) string {
	if format == "" {
		format = DefaultDateTimeFormate
	}
	temp, _ := time.Parse(format, issuedDate)
	return temp.Format(format)
}

func ParseDateTimeDefaultFormate0700(issuedDate string) string {
	temp, _ := time.Parse(DefaultDateTimeFormate0700, issuedDate)
	locAt, _ := time.LoadLocation(DefaultTimeZone)
	return temp.In(locAt).UTC().Format(DefaultDateTimeFormate)
}

// ParseDateTime_CustomFmt is parse date time format string, return value is ( 25 ม.ค.2566 02:02:02)
func ParseDateTimeCustomFmtToShortThai(format, issuedDate string) *string {

	var effectiveDate string

	if format == "" {
		format = "2006-01-02T15:04:05Z07:00"
		// customFmt := "02-01-2006T15:04:05Z07:00"
		// customFmt := "02/01/2006T15:04:05Z07:00"
	}

	pareDateTime, _ := time.Parse(format, issuedDate)
	year, month, day := pareDateTime.Date()
	hour := pareDateTime.Hour()
	minute := pareDateTime.Minute()
	second := pareDateTime.Second()

	effectiveDate = fmt.Sprintf("%v %v%v %02d:%02d:%02d", day, *ParsMonthEnToTh(fmt.Sprintf("%v", month)), year+543, hour, minute, second)

	return &effectiveDate
}

// ParseDateTime_CustomFmt is parse date time format string, return value is ( 25 กุมภาพันธ์ 2566 02:02:02)
func ParseDateTimeCustomFmtToFullThai(format, issuedDate string) *string {

	var effectiveDate string

	if format == "" {
		format = "2006-01-02T15:04:05Z07:00"
		// customFmt := "02-01-2006T15:04:05Z07:00"
		// customFmt := "02/01/2006T15:04:05Z07:00"
	}

	pareDateTime, _ := time.Parse(format, issuedDate)
	year, month, day := pareDateTime.Date()
	hour := pareDateTime.Hour()
	minute := pareDateTime.Minute()
	second := pareDateTime.Second()

	effectiveDate = fmt.Sprintf("%v %v %v %02d:%02d:%02d", day, *ParsMonthEnToThFull(fmt.Sprintf("%v", month)), year+543, hour, minute, second)

	return &effectiveDate
}

func ParsMonthEnToTh(monthEn string) *string {

	var monthTh string

	switch monthEn {
	case "January":
		monthTh = "ม.ค."
	case "February":
		monthTh = "ก.พ."
	case "March":
		monthTh = "มี.ค."
	case "April":
		monthTh = "เม.ย."
	case "May":
		monthTh = "พ.ค."
	case "June":
		monthTh = "มิ.ย."
	case "July":
		monthTh = "ก.ค."
	case "August":
		monthTh = "ส.ค."
	case "September":
		monthTh = "ก.ย."
	case "October":
		monthTh = "ต.ค."
	case "November":
		monthTh = "พ.ย."
	case "December":
		monthTh = "ธ.ค."
	}

	return &monthTh
}

func ParsMonthEnToThFull(monthEn string) *string {

	var monthTh string

	switch monthEn {
	case "January":
		monthTh = "มกราคม"
	case "February":
		monthTh = "กุมภาพันธ์"
	case "March":
		monthTh = "มีนาคม"
	case "April":
		monthTh = "เมษายน"
	case "May":
		monthTh = "พฤษภาคม"
	case "June":
		monthTh = "มิถุนายน"
	case "July":
		monthTh = "กรกฎาคม"
	case "August":
		monthTh = "สิงหาคม"
	case "September":
		monthTh = "กันยายน"
	case "October":
		monthTh = "ตุลาคม"
	case "November":
		monthTh = "พฤศจิกายน"
	case "December":
		monthTh = "ธันวาคม"
	}

	return &monthTh
}

func CalculateStartTimeEndTimeToMillisecondsToInteger(startTime time.Time) int {

	locAt, _ := time.LoadLocation(DefaultTimeZone)
	endTime := time.Now().Local().In(locAt).UTC().Sub(startTime).Milliseconds()
	return int(endTime)

}
