package core

import (
	"github.com/boardware-cloud/model/common"
	"gorm.io/gorm"
)

var db *gorm.DB

var accountRepository AccountRepository
var webauthRepository WebauthRepository

func Init(injectDB *gorm.DB) {
	db = injectDB
	db.AutoMigrate(&Account{},
		&Credential{},
		&SessionData{},
		&Session{},
		&Ticket{},
		&VerificationCode{},
		&ColdDown{})
	accountRepository = NewAccountRepository(db)
	webauthRepository = NewWebauthRepository(db)
	common.Init(injectDB)
}
