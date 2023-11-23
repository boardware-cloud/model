package core

import (
	"github.com/boardware-cloud/model"
)

var db = model.GetDB()

func init() {
	db.AutoMigrate(&Account{},
		&Credential{},
		&SessionData{},
		&Session{},
		&Ticket{},
		&VerificationCode{},
		&ColdDown{})
}
