package core

import (
	"github.com/boardware-cloud/common/constants"
	"github.com/boardware-cloud/common/utils"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID       uint           `gorm:"primarykey"`
	Email    string         `json:"email" gorm:"index:email_index,unique"`
	Password string         `json:"password" gorm:"type:CHAR(128)"`
	Salt     []byte         `json:"salt"`
	Role     constants.Role `json:"role" gorm:"type:VARCHAR(128)"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}
