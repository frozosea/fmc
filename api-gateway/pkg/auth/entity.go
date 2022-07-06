package auth

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserWithId struct {
	Id int
	User
}
