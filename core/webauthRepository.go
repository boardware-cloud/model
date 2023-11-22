package core

import (
	"github.com/boardware-cloud/model"
	"gorm.io/gorm"
)

var webauthRepository *WebauthRepository

func GetWebauthRepository() *WebauthRepository {
	if webauthRepository == nil {
		webauthRepository = NewWebauthRepository()
	}
	return webauthRepository
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
