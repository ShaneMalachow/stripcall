package stripcall

import "github.com/jinzhu/gorm"

/*
After adding the model here, add it to the Automigrate() call in database.go
*/

type User struct {
	gorm.Model
	firstName string
	lastName  string
}

type Event struct {
	gorm.Model
	name     string
	year     int
	headTech User
}
