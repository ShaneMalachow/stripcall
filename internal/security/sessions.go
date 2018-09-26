package security

import (
	"github.com/jinzhu/gorm"
	"github.com/shanemalachow/stripcall/internal/user"
	"time"
)

type Session struct {
	gorm.Model
	user.User
	UserId     uint
	token      uint64
	expiration time.Time
}

type SessionService struct {
	gorm.DB
}

func (svc SessionService) CreateSession(user user.User)
