package security

import (
	"github.com/jinzhu/gorm"
	"github.com/shanemalachow/stripcall/internal/user"
	"time"
)

//Session represents a user session on the application
type Session struct {
	gorm.Model
	user.User
	UserID     uint
	token      uint64
	expiration time.Time
}

//SessionService allows access to CRUD operations for Sessions
type SessionService struct {
	gorm.DB
}

//CreateSession creates a new Session for a user
func (svc SessionService) CreateSession(user user.User) {

}
