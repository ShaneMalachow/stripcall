package security

import (
	"github.com/jinzhu/gorm"
	"github.com/shanemalachow/stripcall/internal/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Auth struct {
	gorm.Model
	user.User
	UserId       uint
	PasswordHash []byte
}

type AuthService struct {
	gorm.DB
}

func (auth Auth) CheckPassword(password []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(auth.PasswordHash, password)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

func HashNewPassword(password []byte) ([]byte, error) {
	saltAndHash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return make([]byte, 0), err
	}
	return saltAndHash, nil
}
