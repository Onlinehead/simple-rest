package birthday

import (
	"log"
	"time"
)

func GetDaysBeforeBirthday(birthdayDate time.Time, now time.Time) int {
	curYear := now.Year()
	bMonth, bDay := FixLeapYearDate(curYear, int(birthdayDate.Month()), birthdayDate.Day())
	birthdayCurrentYear, err := GetTimeObjFromVals(curYear, bMonth, bDay)
	if err != nil {
		log.Fatalln("Cannot parse date:", err)
	}
	daysBefore := DaysDeltaBetweenTimes(birthdayCurrentYear, now)
	if daysBefore < 0 {
		bMonth, bDay = FixLeapYearDate(curYear+1, int(birthdayDate.Month()), birthdayDate.Day())
		birthdayNextYear, err := GetTimeObjFromVals(curYear+1, int(birthdayDate.Month()),  birthdayDate.Day())
		if err != nil {
			log.Fatalln("Cannot parse date:", err)
		}
		daysBefore = DaysDeltaBetweenTimes(birthdayNextYear, now)
	}
	return daysBefore
}

func IsBirthdayInFuture(birthday time.Time, now time.Time) bool {
	if now.Unix() < birthday.Unix(){
		return false
	}
	return true
}