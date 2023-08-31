package core

import "gorm.io/gorm"

var db *gorm.DB

func Init(injectDB *gorm.DB) {
	db = injectDB
	db.AutoMigrate(&Credential{})
	db.AutoMigrate(&LoginRecord{})
	db.AutoMigrate(&SessionData{})
	db.AutoMigrate(Session{})
	db.AutoMigrate(&Ticket{})
	db.AutoMigrate(&VerificationCode{})
	db.AutoMigrate(&Account{})
}
