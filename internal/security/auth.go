package security

import (
	"github.com/jinzhu/gorm"
	"github.com/shanemalachow/stripcall/internal/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

//Auth representing a password hash and user account mashup
type Auth struct {
	gorm.Model
	user.User
	UserID       uint
	PasswordHash []byte
}

//AuthService provides a way to access DB and CRUD Auth objects
type AuthService struct {
	gorm.DB
}

//CheckPassword checks whether a password matches a provided salted hash
func (auth Auth) CheckPassword(password []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(auth.PasswordHash, password)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

//HashNewPassword creates a new salted hash based on an input password
func HashNewPassword(password []byte) ([]byte, error) {
	saltAndHash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return make([]byte, 0), err
	}
	return saltAndHash, nil
}
