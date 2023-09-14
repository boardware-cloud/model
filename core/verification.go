package core

import (
	constants "github.com/boardware-cloud/common/constants/account"
	"github.com/boardware-cloud/common/utils"
	"gorm.io/gorm"
)

type VerificationCode struct {
	gorm.Model
	Identity string                            `json:"email" gorm:"index:verification_index"` // Email or phone number
	Purpose  constants.VerificationCodePurpose `json:"purpose" gorm:"type:VARCHAR(128)"`
	Code     string                            `json:"code" gorm:"type:CHAR(6)"`
	Tries    int64
}

func (v *VerificationCode) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = utils.GenerteId()
	return
}
