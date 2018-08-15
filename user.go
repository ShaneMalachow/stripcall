package main

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
	Fencer
)

type UserService interface {
	User(id int) (*User, error)
	UsersByRole(role Role) ([]*User, error)
	CreateUser(user *User) error
	DeleteUser(id int) error
}
