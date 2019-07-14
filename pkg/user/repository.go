package user

import "time"

type Repository interface {
	FindUser(username string) (*User, error)
	AddUser(username string, birthday time.Time) error
	Close()
}
