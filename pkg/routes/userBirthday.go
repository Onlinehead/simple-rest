package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/onlinehead/simple-rest/pkg/birthday"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)


func UserBirthday(c *gin.Context) {
	username := c.Param("username")
	if ! IsLetter(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must contains only letters"})
		log.Errorln("Username contains not only letters: ", username)
		return
	}
	user, err := Repo.FindUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to get info about user %v", username)})
		log.Errorf("Error durring getting info for user '%v': %v", username, err)
		return
	}
	daysBeforeBirthday := birthday.GetDaysBeforeBirthday(time.Unix(user.Birthday, 0), time.Now())
	if daysBeforeBirthday == 0 {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %v! Happy birthday!", username)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %v! Your birthday in %v day(s)", username, daysBeforeBirthday)})
	}
}
