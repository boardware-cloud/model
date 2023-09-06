package core

import (
	"github.com/boardware-cloud/model/common"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(injectDB *gorm.DB) {
	db = injectDB
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Credential{})
	db.AutoMigrate(&LoginRecord{})
	db.AutoMigrate(&SessionData{})
	db.AutoMigrate(Session{})
	db.AutoMigrate(&Ticket{})
	db.AutoMigrate(&VerificationCode{})
	common.Init(injectDB)
}
