package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/shanemalachow/stripcall/internal/security"
	"github.com/shanemalachow/stripcall/internal/user"
	"log"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	db.AutoMigrate(&user.User{}, &security.Auth{})
	testUser := user.User{FirstName: "Shane", LastName: "Malachow", Armorer: true, Admin: true}
	fmt.Println(testUser)
	db.Create(&testUser)

	var shane user.User
	db.First(&shane, "FirstName = ?", "Shane")
	fmt.Println(shane.FirstName + shane.LastName)
}
