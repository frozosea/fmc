package auth

import "golang.org/x/crypto/bcrypt"

type IHash interface {
	Hash(pwd string) (string, error)
	CheckPasswordHash(hashedPassword string, saltPassword string) bool
}
type Hash struct {
}

func (h *Hash) Hash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}
func (h *Hash) CheckPasswordHash(originPassword string, saltPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(originPassword), []byte(saltPassword))
	return err == nil
}
