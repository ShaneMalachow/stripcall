package stripcall

import (
	"github.com/jinzhu/gorm"
	"time"
)

/*
After adding the model here, add it to the Automigrate() call in database.go
*/

type User struct {
	gorm.Model    `json:"-"`
	UserName      string `gorm:"PRIMARY_KEY" json:"username"`
	FirstName     string `gorm:"NOT NULL" json:"firstName"`
	LastName      string `gorm:"NOT NULL" json:"lastName"`
	PhoneNumber   string `json:"phoneNumber"`
	Armorer       bool   `gorm:"DEFAULT:false" json:"armorer"`
	RefereeFoil   rune   `gorm:"DEFAULT:0" json:"refFoil"`
	RefereeEpee   rune   `gorm:"DEFAULT:0" json:"refEpee"`
	RefereeSaber  rune   `gorm:"DEFAULT:0" json:"refSaber"`
	BoutCommittee bool   `gorm:"DEFAULT:false" json:"bc"`
	Admin         bool   `gorm:"DEFAULT:false" json:"admin"`
}

type Event struct {
	gorm.Model `json:"-"`
	Name       string `gorm:"NOT NULL" json:"name"`
	Location   string `json:"location"`
	Month      int    `gorm:"NOT NULL, DEFAULT:0" json:"month"`
	Year       int    `gorm:"NOT NULL, DEFAULT:0" json:"year"`
	HeadTech   User   `json:"-"`
	HeadTechID string `json:"headTech"`
	HeadRef    User   `json:"-"`
	HeadRefID  string `json:"headRef"`
	HeadBC     User   `json:"-"`
	HeadBCID   string `json:"headBC"`
}

type Call struct {
	gorm.Model  `json:"-"`
	Reporter    User      `gorm:"NOT NULL" json:"-"`
	ReporterID  string    `json:"reporter"`
	Responder   User      `json:"-"`
	ResponderID string    `json:"responder"`
	Issue       string    `gorm:"NOT NULL" json:"issue"`
	Strip       string    `gorm:"NOT NULL" json:"strip"`
	ReportTime  time.Time `gorm:"NOT NULL, DEFAULT:current_time" json:"time"`
}

type Message struct {
	gorm.Model `json:"-"`
	Call       Call
	CallID     Call
	Sender     User
	Message    string
	StaffOnly  bool
}

func current_time() time.Time {
	return time.Now()
}
