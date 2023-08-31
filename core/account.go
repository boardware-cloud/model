package core

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/boardware-cloud/common/constants"
	"github.com/boardware-cloud/common/utils"
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
	Role               constants.Role `json:"role" gorm:"type:VARCHAR(128)"`
	Totp               *string
	WebAuthnCredential []Credential
	WebAuthnSession    []SessionData
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
		return credential.Credential
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
	WebAuthnSessionData
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
	webauthn.Credential
	AccountId uint
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}
