package stripcall

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
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
		return nil
	}

	db.AutoMigrate(
		&User{},
		&Event{},
		&Call{},
		&Message{},
	)

	return db
}

func LookupAndWriteJSON(db *gorm.DB, w *http.ResponseWriter, key interface{}, object interface{}) {
	db.Where(key).Find(&object)
	err := json.NewEncoder(*w).Encode(object)
	if err != nil {
		HandleError(w, err.Error(), http.StatusInternalServerError)
	}
}

func LookupAndWriteAllJSON(db *gorm.DB, w *http.ResponseWriter, objects interface{}) {
	db.Find(&objects)
	err := json.NewEncoder(*w).Encode(objects)
	if err != nil {
		HandleError(w, err.Error(), http.StatusInternalServerError)
	}
}
