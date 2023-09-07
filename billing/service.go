package billing

import (
	"github.com/boardware-cloud/common/constants"
	"github.com/boardware-cloud/common/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(injectDB *gorm.DB) {
	db = injectDB
}

type Service struct {
	gorm.Model
	Name        constants.ServiceName `json:"name" gorm:"index:name_index,unique"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Url         string                `json:"url"`
	Type        constants.ServiceType `json:"type" gorm:"type:VARCHAR(128)"`
}

func (a *Service) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}

func AutoMigrate(name constants.ServiceName, title, description, url string, serviceType constants.ServiceType) {
	service := Service{}
	db.Where("name = ?", name).Find(&service)
	service.Title = title
	service.Description = description
	service.Url = url
	service.Type = serviceType
	db.Save(&service)
}

// func NewService(name constants.ServiceName, title, description, url string, serviceType constants.ServiceType) Service {
// 	service := Service{
// 		Name:        name,
// 		Title:       title,
// 		Description: description,
// 		Url:         url,
// 		Type:        serviceType,
// 	}
// 	service.ID = utils.GenerteId()
// 	return service
// }
