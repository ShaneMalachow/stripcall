package user

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Role      Role
}

type Role int

const (
	Armorer Role = iota
	Referee
	BoutCommittee
	Medical
	Replay
	Fencer
)

type UserService interface {
	User(id int) (*User, error)
	UsersByRole(role Role) ([]*User, error)
	FindUser(fn string, ln string) (*User, error)
	NewUser(user *User) error
	DeleteUser(id int) error
}

func (user *User) FullName() (string, error) {
	if user.FirstName == "" && user.LastName == "" {
		return "", errors.New("user name is empty")
	}
	return user.FirstName + " " + user.LastName, nil
}

func CreateUser(fn string, ln string, role Role) (*User, error) {
	return &User{FirstName: fn, LastName: ln, Role: role}, nil
}
