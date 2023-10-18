package core

import (
	"time"

	constants "github.com/boardware-cloud/common/constants/account"
	"github.com/boardware-cloud/common/utils"
	"github.com/boardware-cloud/model/abstract"
	"github.com/boardware-cloud/model/common"
	"github.com/chenyunda218/golambda"
	"github.com/go-webauthn/webauthn/webauthn"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email              string         `json:"email" gorm:"index:email_index,unique"`
	Password           string         `json:"-" gorm:"type:CHAR(128)"`
	Salt               []byte         `json:"salt"`
	Role               constants.Role `json:"role"`
	Totp               *string
	WebAuthnCredential []Credential
	WebAuthnSession    []SessionData
}

func (Account) Owner() abstract.Owner {
	return nil
}

func (a Account) ID() uint {
	return a.Model.ID
}

func (Account) TypeName() string {
	return "account"
}

func (a Account) Own(asset abstract.Asset) bool {
	return abstract.Own(a, asset)
}

func ListAccount(index, limit int64) common.List[Account] {
	return common.ListModel(&[]Account{}, index, limit)
}

func (a Account) CreateColdDown() {
	db.Save(&ColdDown{AccountId: a.ID()})
}

func (a Account) ColdDown(coldDownTime int64) bool {
	var coldDown ColdDown
	db.Where("account_id = ?", a.ID).Order("created_at DESC").Limit(1).Find(&coldDown)
	return time.Now().UnixMilli()-coldDown.CreatedAt.UnixMilli() > coldDownTime
}

func (a Account) WebAuthnID() []byte {
	return []byte(utils.UintToString(a.ID()))
}

func (a Account) WebAuthnName() string {
	return a.Email
}

func (a Account) WebAuthnDisplayName() string {
	return a.Email
}

func (Account) WebAuthnIcon() string {
	return ""
}

func (a Account) WebAuthnCredentials() []webauthn.Credential {
	var credentials []Credential = make([]Credential, 0)
	db.Where("account_id = ?", a.ID).Find(&credentials)
	return golambda.Map(credentials, func(_ int, credential Credential) webauthn.Credential {
		return webauthn.Credential(credential.Credential)
	})
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = utils.GenerteId()
	return err
}

func (a *Account) SetPassword(password string) *Account {
	hashed, salt := utils.HashWithSalt(password)
	a.Password = hashed
	a.Salt = salt
	return a
}

type ColdDown struct {
	gorm.Model
	AccountId uint `json:"email" gorm:"index:account_id_index"`
}

func (a *ColdDown) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return err
}
