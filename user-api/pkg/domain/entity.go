package domain

import "user-api/pkg/user"

type Container struct {
	Number string `json:"number"`
}

type UserWithId struct {
	Id int
	user.User
}
