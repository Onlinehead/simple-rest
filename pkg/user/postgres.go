package user

import (
	"errors"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
	"time"
)

type repo struct {
	conn *pg.DB
}

var UsernameRecordDuplicateErr = errors.New(
	"ERROR #23505 duplicate key value violates unique constraint \"users_pkey\"")

func NewPostgresRepo(addr string, user string, password string, database string, migrationsDir string) (Repository, error) {
	db := pg.Connect(&pg.Options{
		Addr: addr,
		User: user,
		Password: password,
		Database: database,
	})
	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	col := migrations.NewCollection()
	err = col.DiscoverSQLMigrations(migrationsDir)
	if err != nil {
		return nil, err
	}
	col.Run(db,"init")
	oldVersion, newVersion, err := col.Run(db,"up")
	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		log.Printf("DB version is %d\n", oldVersion)
	}
	if err != nil {
		return nil, err
	}
	return &repo{
		conn: db,
	}, nil
}

func (r *repo) FindUser(username string) (*User, error) {
	user := &User {
		Username: username,
	}
	err := r.conn.Select(user)
	return user, err
}

func (r *repo) AddUser(username string, birthday time.Time) (err error) {
	user := &User{
		Username: username,
		Birthday: birthday.Unix(),
	}
	log.Debug(user)
	_, err = r.conn.Model(user).Insert(user)
	if err != nil {
		if err.Error() == UsernameRecordDuplicateErr.Error() {
			return nil
		}
		return
	}

	return
}

func (r *repo) Close () {
	r.conn.Close()
}