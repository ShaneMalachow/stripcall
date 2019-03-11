package stripcall

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type DependencyMap struct {
	DB   *gorm.DB
	Conf map[string]string
}

type PostResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Setup(r *mux.Router, dep *DependencyMap) *http.Server {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	ConfigRouter(apiRouter, dep)
	return &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func ParseConfig(confLoc string) map[string]string {
	return map[string]string{
		"dbType":    "sqlite3",
		"dbConnect": "./test.sqlite",
	}
}
