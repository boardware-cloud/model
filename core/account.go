package core

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/boardware-cloud/common/code"
	constants "github.com/boardware-cloud/common/constants/account"
	"github.com/boardware-cloud/common/utils"
	"github.com/boardware-cloud/model/common"
	"github.com/chenyunda218/golambda"
	"github.com/go-webauthn/webauthn/webauthn"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID                 uint           `gorm:"primarykey"`
	Email              string         `json:"email" gorm:"index:email_index,unique"`
	Password           string         `json:"password" gorm:"type:CHAR(128)"`
	Salt               []byte         `json:"salt"`
	Role               constants.Role `json:"role"`
	Totp               *string
	WebAuthnCredential []Credential
	WebAuthnSession    []SessionData
}

func FindAccount(conds ...any) (Account, error) {
	var account Account
	if ctx := db.Find(&account, conds...); ctx.RowsAffected == 0 {
		return account, code.ErrNotFound
	}
	return account, nil
}

func FindAccountById(id uint) (Account, error) {
	return FindAccount(id)
}

func FindAccountByEmail(email string) (Account, error) {
	return FindAccount("email = ?", email)
}

func ListAccount(index, limit int64) common.List[Account] {
	var accounts []Account
	return common.ListModel(&accounts, index, limit)
}

func (a Account) CreateColdDown() {
	CreateColdDown(a.ID)
}

func (a Account) ColdDown(coldDownTime int64) bool {
	var coldDown ColdDown
	db.Where("account_id = ?", a.ID).Order("created_at DESC").Limit(1).Find(&coldDown)
	return time.Now().UnixMilli()-coldDown.CreatedAt.UnixMilli() > coldDownTime
}

func (a Account) WebAuthnID() []byte {
	return []byte(utils.UintToString(a.ID))
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

type WebAuthnSessionData webauthn.SessionData

func (w *WebAuthnSessionData) Scan(value any) error {
	return json.Unmarshal(value.([]byte), w)
}

func (w WebAuthnSessionData) Value() (driver.Value, error) {
	b, err := json.Marshal(w)
	return b, err
}

func (WebAuthnSessionData) GormDataType() string {
	return "JSON"
}

type SessionData struct {
	gorm.Model
	AccountId uint
	Data      WebAuthnSessionData
}

func (s *SessionData) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = utils.GenerteId()
	return
}

type WebAuthnCredential webauthn.Credential

func (w *WebAuthnCredential) Scan(value any) error {
	return json.Unmarshal(value.([]byte), w)
}

func (w WebAuthnCredential) Value() (driver.Value, error) {
	b, err := json.Marshal(w)
	return b, err
}

func (WebAuthnCredential) GormDataType() string {
	return "JSON"
}

type Credential struct {
	gorm.Model
	Name       string
	Os         string
	Credential WebAuthnCredential
	AccountId  uint
}

func (a *Credential) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return err
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return err
}

type ColdDown struct {
	gorm.Model
	AccountId uint `json:"email" gorm:"index:account_id_index"`
}

func (a *ColdDown) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return err
}

func CreateColdDown(accountId uint) {
	db.Save(&ColdDown{AccountId: accountId})
}
