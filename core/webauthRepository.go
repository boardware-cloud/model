package core

import (
	"github.com/Dparty/common/singleton"
	"github.com/boardware-cloud/model"
	"gorm.io/gorm"
)

var webauthRepository = singleton.NewSingleton[WebauthRepository](NewWebauthRepository, singleton.Eager)

func GetWebauthRepository() *WebauthRepository {
	return webauthRepository.Get()
}

func NewWebauthRepository() *WebauthRepository {
	return &WebauthRepository{model.GetDB()}
}

type WebauthRepository struct {
	db *gorm.DB
}

func (w WebauthRepository) List(conds ...any) []Credential {
	var credentials []Credential
	w.db.Find(&credentials, conds...)
	return credentials
}

var sessionDataRepository *SessionDataRepository

func GetSessionDataRepository() *SessionDataRepository {
	if sessionDataRepository == nil {
		sessionDataRepository = NewSessionDataRepository()
	}
	return sessionDataRepository
}

func NewSessionDataRepository() *SessionDataRepository {
	return &SessionDataRepository{model.GetDB()}
}

type SessionDataRepository struct {
	db *gorm.DB
}

func (s SessionDataRepository) Find(conds ...any) *SessionData {
	var sessionData SessionData
	ctx := s.db.Find(&sessionData, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &sessionData
}

func (s SessionDataRepository) GetById(id uint) *SessionData {
	return s.Find(id)
}

func (s SessionDataRepository) Save(sessionData *SessionData) {
	s.db.Save(sessionData)
}
