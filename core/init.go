package core

import (
	"github.com/boardware-cloud/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = model.GetDB()
	db.AutoMigrate(&Account{},
		&Credential{},
		&SessionData{},
		&Session{},
		&Ticket{},
		&VerificationCode{},
		&ColdDown{})
}
