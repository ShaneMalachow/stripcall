package stripcall

import "github.com/jinzhu/gorm"

func Connect(dbType string, connector string) *gorm.DB {
	var db *gorm.DB
	var err error
	switch dbType {
	case "sqlite3":
		db, err = gorm.Open("sqlite3", connector)
	}

	if err != nil {
		return nil
	}

	db.AutoMigrate(
		&User{},
		&Event{},
	)

	return db
}
