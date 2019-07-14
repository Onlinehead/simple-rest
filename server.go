package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/onlinehead/simple-rest/pkg/routes"
	"github.com/onlinehead/simple-rest/pkg/user"
	log "github.com/sirupsen/logrus"
	"os"
)
var (
	AppVer string
	BuildTime string
)

type Configuration struct {
	DatabaseType string `default:"postgres"`
	PostgresHost string `default:"localhost"`
	PostgresPort string `default:"5432"`
	PostgresDb string `default:"simple_rest"`
	PostgresUser	string `default:"user"`
	PostgresPassword	string `default:"zzz"`
	Interface string `default:"0.0.0.0"`
	Port string `default:"8080"`
}

func initLogger() {
	log.SetOutput(os.Stdout)
}

func initPostgresDB(addr string, username string, password string, database string, migrationsDir string) (user.Repository, error) {
	repo, err := user.NewPostgresRepo(addr, username, password, database, migrationsDir)
	return repo, err
}

func SetupGin() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", routes.Ping)
	router.PUT("/hello/:username", routes.AddUser)
	router.GET("/hello/:username", routes.UserBirthday)
	return router
}

func main() {
	var s Configuration
	err := envconfig.Process("rest", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	initLogger()
	log.Info("App version: ", AppVer)
	log.Info("Build time: ", BuildTime)
	if s.DatabaseType != "postgres" {
		log.Fatalln("Database type not supported: ", s.DatabaseType)
	}
	repo, err := initPostgresDB(fmt.Sprintf("%v:%v", s.PostgresHost, s.PostgresPort),s.PostgresUser, s.PostgresPassword, s.PostgresDb, "postgres_migrations")
	if err != nil {
		log.Fatalln("Unexpected error happened during database connection init:", err)
	}
	routes.Repo = repo
	defer repo.Close()
	router := SetupGin()
	err = router.Run(fmt.Sprintf("%v:%v", s.Interface, s.Port))
	if err != nil {
		log.Fatalln("Unable to start HTTP server:", err)
	}
}
