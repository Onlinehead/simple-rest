package birthday

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetDaysBeforeBirthday(t *testing.T) {
	date1, _ := ParseDate("2017-06-19")
	date2, _ := ParseDate("2017-06-22")
	assert.Equal(t, GetDaysBeforeBirthday(date1, date2), 362)
	assert.Equal(t, GetDaysBeforeBirthday(date2, date1), 3)
}

func TestIsBirthdayInFuture(t *testing.T) {
	tomorrow := time.Now().AddDate(0,0,1)
	now := time.Now()
	assert.Equal(t, IsBirthdayInFuture(tomorrow, now), false)
	assert.Equal(t, IsBirthdayInFuture(now, tomorrow), true)
}