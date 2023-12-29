package core

import (
	"github.com/Dparty/common/singleton"
	constants "github.com/boardware-cloud/common/constants/account"
	"github.com/boardware-cloud/model"
	"gorm.io/gorm"
)

var verificationCodeRepository = singleton.NewSingleton[VerificationCodeRepository](NewVerificationCodeRepository, singleton.Eager)

func GetVerificationCodeRepository() *VerificationCodeRepository {
	return verificationCodeRepository.Get()
}

type VerificationCodeRepository struct {
	db *gorm.DB
}

func NewVerificationCodeRepository() *VerificationCodeRepository {
	return &VerificationCodeRepository{model.GetDB()}
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
