package stripcall

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

type DependencyMap struct {
	DB   *gorm.DB
	Conf map[string]string
}

func Setup(r *mux.Router, dep DependencyMap) *http.Server {

}

func ParseConfig(confLoc string) map[string]string {

}
