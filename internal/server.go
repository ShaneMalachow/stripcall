package stripcall

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
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
		Addr:         os.Getenv("STRIPCALL_ADDR") + ":" + os.Getenv("STRIPCALL_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func ParseConfig(confLoc string) map[string]string {
	return map[string]string{
		"dbType":    os.Getenv("DBTYPE"),
		"dbConnect": os.Getenv("DBCONNECTOR"),
	}
}
