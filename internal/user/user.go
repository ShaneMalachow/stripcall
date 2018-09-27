package user

import (
	"errors"
	"github.com/jinzhu/gorm"
)

//User represents an application user
type User struct {
	gorm.Model
	FirstName       string
	LastName        string
	Admin           bool
	Armorer         bool
	Medical         bool
	Referee         bool
	BoutCommittee   bool
	HeadTech        bool
	RefereeAssigner bool
}

//UserService allows access to CRUD operations for User objects
type Service interface {
	User(id int) (*User, error)
	UsersByRole(role Role) ([]*User, error)
	FindUser(fn string, ln string) (*User, error)
	NewUser(user *User) error
	DeleteUser(id int) error
}

//FullName returns the full name of the user as one string
func (user *User) FullName() (string, error) {
	if user.FirstName == "" && user.LastName == "" {
		return "", errors.New("user name is empty")
	}
	return user.FirstName + " " + user.LastName, nil
}
