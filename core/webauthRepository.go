package core

import "gorm.io/gorm"

func NewWebauthRepository(db *gorm.DB) WebauthRepository {
	return WebauthRepository{db}
}

type WebauthRepository struct {
	db *gorm.DB
}

func (w WebauthRepository) List(conds ...any) []Credential {
	var credentials []Credential
	w.db.Find(&credentials, conds...)
	return credentials
}
