package core

import (
	constants "github.com/boardware-cloud/common/constants/account"
	"gorm.io/gorm"
)

type VerificationCodeRepository struct {
	db *gorm.DB
}

func (v VerificationCodeRepository) Find(conds ...any) *VerificationCode {
	var verificationCode VerificationCode
	ctx := v.db.Find(&verificationCode, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &verificationCode
}

func (v VerificationCodeRepository) Save(code *VerificationCode) {
	v.db.Save(code)
}

func (v VerificationCodeRepository) Get(email string, purpose constants.VerificationCodePurpose) *VerificationCode {
	var verificationCode VerificationCode
	ctx := v.db.Where("identity = ?", email).Where("purpose = ?", purpose).Order("created_at DESC").Find(&verificationCode)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &verificationCode
}

func (v VerificationCodeRepository) Delete(email string, purpose constants.VerificationCodePurpose) {
	v.db.Where("identity = ?", email).Where("purpose = ?", purpose).Delete(&VerificationCode{})
}
