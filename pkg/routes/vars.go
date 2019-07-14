package routes

import "github.com/onlinehead/simple-rest/pkg/user"

var Repo user.Repository

type UserPayload struct {
	DateOfBirth string `json:"dateOfBirth" binding:"required"`
}