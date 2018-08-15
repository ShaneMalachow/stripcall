package user

import "errors"

type User struct {
	id        int
	firstName string
	lastName  string
	role      Role
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
	if user.firstName == "" && user.lastName == "" {
		return "", errors.New("user name is empty")
	}
	return user.firstName + " " + user.lastName, nil
}

func CreateUser(id int, fn string, ln string, role Role) (*User, error) {
	return &User{id: id, firstName: fn, lastName: ln, role: role}, nil
}
