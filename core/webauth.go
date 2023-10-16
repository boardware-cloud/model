package core

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/boardware-cloud/common/utils"
	"github.com/go-webauthn/webauthn/webauthn"
	"gorm.io/gorm"
)

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
