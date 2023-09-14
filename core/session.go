package core

import (
	"time"

	constants "github.com/boardware-cloud/common/constants/account"
	"github.com/boardware-cloud/common/utils"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	AccountId   uint                  `json:"accountId" gorm:"index:accountId_index"`
	Token       string                `json:"token"`
	TokeType    constants.TokenType   `json:"tokenType" gorm:"type:VARCHAR(128)"`
	TokenFormat constants.TokenFormat `json:"tokenFormat" gorm:"type:VARCHAR(128)"`
	ExpiredAt   time.Time             `json:"expiredAt"`
}

func (a *Session) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}
