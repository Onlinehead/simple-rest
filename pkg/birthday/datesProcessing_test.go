package birthday

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestIsLeapYear(t *testing.T) {
	assert.Equal(t, IsLeapYear(2007), false)
	assert.Equal(t, IsLeapYear(2020), true)
}

func TestParseDate(t *testing.T) {
	_, err := ParseDate("2017-06-19")
	assert.Equal(t, err, nil)
	_, err = ParseDate("2017-13-19")
	assert.NotEqual(t, err, nil)
	_, err = ParseDate("13-02-2017")
	assert.NotEqual(t, err, nil)
	_, err = ParseDate("13-02-2017-123-432")
	assert.NotEqual(t, err, nil)
	_, err = ParseDate("not a date")
	assert.NotEqual(t, err, nil)
}

func TestFixLeapYearDate(t *testing.T) {
	// 2020 is a leap year
	month, day := FixLeapYearDate(2020, 02,29)
	assert.Equal(t, month, 02)
	assert.Equal(t, day, 29)
	month, day = FixLeapYearDate(2020, 02,27)
	assert.Equal(t, month, 02)
	assert.Equal(t, day, 27)
	// 2018 is NOT a leap year
	month, day = FixLeapYearDate(2019, 02,29)
	assert.Equal(t, month, 03)
	assert.Equal(t, day, 01)
	month, day = FixLeapYearDate(2019, 02,27)
	assert.Equal(t, month, 02)
	assert.Equal(t, day, 27)
}

func TestGetTimeObjFromVals(t *testing.T) {
	_, err := GetTimeObjFromVals(2018, 05, 12)
	assert.Equal(t, err, nil)
	_, err = GetTimeObjFromVals(2018, 05, 32)
	assert.NotEqual(t, err, nil)
	_, err = GetTimeObjFromVals(-2018, 05, 32)
	assert.NotEqual(t, err, nil)
}

func TestDaysDeltaBetweenTimes(t *testing.T) {
	from, _ := GetTimeObjFromVals(2018, 05, 01)
	to, _ := GetTimeObjFromVals(2018, 05, 01)
	assert.Equal(t, DaysDeltaBetweenTimes(from, to), 0)

	from, _ = GetTimeObjFromVals(2018, 05, 01)
	to, _ = GetTimeObjFromVals(2018, 05, 02)
	assert.Equal(t, DaysDeltaBetweenTimes(to, from), 1)
	assert.Equal(t, DaysDeltaBetweenTimes(from, to), -1)

	tomorrowMidnight := time.Now().AddDate(0,0,1)
	now := time.Now().Add(5 * time.Minute)
	assert.Equal(t, DaysDeltaBetweenTimes(tomorrowMidnight, now), 1)
}