package core

import (
	"github.com/boardware-cloud/common/code"
	constants "github.com/boardware-cloud/common/constants/account"
	"github.com/boardware-cloud/common/utils"
	"github.com/boardware-cloud/model"
	"gorm.io/gorm"
)

var accountRepository *AccountRepository

func GetAccountRepository() *AccountRepository {
	if accountRepository == nil {
		accountRepository = NewAccountRepository()
	}
	return accountRepository
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{model.GetDB()}
}

type AccountRepository struct {
	db *gorm.DB
}

func (a AccountRepository) Find(conds ...any) *Account {
	var account Account
	ctx := a.db.Find(&account, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &account
}

func (a AccountRepository) GetById(id uint) *Account {
	return a.Find(id)
}

func (a AccountRepository) GetByEmail(email string) *Account {
	return a.Find("email = ?", email)
}

func (a AccountRepository) Create(email, password string, role constants.Role) (*Account, error) {
	if a.GetByEmail(email) != nil {
		return nil, code.ErrEmailExists
	}
	hashed, salt := utils.HashWithSalt(password)
	account := Account{Email: email, Role: role, Password: hashed, Salt: salt}
	switch role {
	case constants.ROOT, constants.ADMIN, constants.USER:
		account.Role = role
	default:
		account.Role = constants.USER
	}
	return &account, nil
}

func (a AccountRepository) Save(account *Account) {
	a.db.Save(account)
}

func NewVerificationCodeRepository() *VerificationCodeRepository {
	return &VerificationCodeRepository{model.GetDB()}
}
