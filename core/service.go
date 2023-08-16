package core

import (
	"github.com/boardware-cloud/common/constants"

	"github.com/boardware-cloud/common/utils"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string                `json:"name" gorm:"index:name_index,unique"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Url         string                `json:"url"`
	Type        constants.ServiceType `json:"type" gorm:"type:VARCHAR(128)"`
}

func (a *Service) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}

func NewService(name, title, description, url string, serviceType constants.ServiceType) Service {
	service := Service{
		Name:        name,
		Title:       title,
		Description: description,
		Url:         url,
		Type:        serviceType,
	}
	service.ID = utils.GenerteId()
	return service
}
