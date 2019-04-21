package stripcall

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect(dbType string, connector string) *gorm.DB {
	var db *gorm.DB
	var err error
	switch dbType {
	case "sqlite3":
		db, err = gorm.Open("sqlite3", connector)
	case "postgres":
		db, err = gorm.Open("postgres", connector)
	}

	if err != nil {
		fmt.Println(err.Error())
		panic("BAD DATABASE CONNECTION")
	}

	db.AutoMigrate(
		&User{},
		&Event{},
		&Call{},
		&Message{},
	)

	return db
}
