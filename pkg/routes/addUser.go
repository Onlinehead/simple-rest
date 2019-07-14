package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onlinehead/simple-rest/pkg/birthday"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)


func AddUser(c *gin.Context) {
	username := c.Param("username")
	if ! IsLetter(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must contains only letters"})
		log.Errorln("Username contains not only letters: ", username)
		return
	}
	var json UserPayload
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Unable to unmarshal JSON"})
		log.Errorln("Unable to unmarshal JSON: ", err)
		return
	}
	birthdayPayload, err := birthday.ParseDate(json.DateOfBirth)
	if err != nil {
		errMess := "Unable to parse a JSON payload"
		if strings.Contains(err.Error(), "day out of range") {
			errMess = "Date of Birth is an incorrect date"
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errMess})
		log.Errorln("Unable to parse a JSON payload: ", err)
		return
	}
	if ! birthday.IsBirthdayInFuture(birthdayPayload, time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to validate birthday date"})
		log.Errorln("Unable to parse a JSON payload: ", err)
		return
	}
	err = Repo.AddUser(username, birthdayPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save data into DB"})
		log.Errorln("Unable to save data into DB: ", err)
		return
	}
	c.String(http.StatusNoContent, "")
}