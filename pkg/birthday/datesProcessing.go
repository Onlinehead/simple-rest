package birthday

import (
	"fmt"
	"time"
)

const (
	dateFormat = "2006-01-02"
)


func ParseDate(date string) (time.Time, error) {
	return time.Parse(dateFormat, date)
}

func IsLeapYear(y int) bool {
	year := time.Date(y, time.December, 31, 0, 0, 0, 0, time.Local)
	days := year.YearDay()

	if days > 365 {
		return true
	} else {
		return false
	}
}

func FixLeapYearDate(year int, month int, day int) (int, int) {
	if day == 29 && month == 2 && !IsLeapYear(year) {
		month = 3
		day = 1
	}
	return month, day
}

func GetTimeObjFromVals(year int, month int, day int) (time.Time, error){
	bCurYear := fmt.Sprintf("%v-%02d-%02d", year, month, day)
	timeObj, err := ParseDate(bCurYear)
	return timeObj, err
}

func DaysDeltaBetweenTimes(from time.Time, to time.Time) int {
	delta := from.Sub(to)
	daysDelta := delta.Hours()/24
	if from.Month() == to.Month() && from.Day() == to.Day() {
		return 0
	} else if daysDelta > 1 || daysDelta < 0{
		return int(daysDelta)
	} else {
		return 1
	}
}