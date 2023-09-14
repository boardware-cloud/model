package billing

import (
	constants "github.com/boardware-cloud/common/constants/service"
	"github.com/boardware-cloud/common/utils"
	"github.com/boardware-cloud/model/common"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(injectDB *gorm.DB) {
	db = injectDB
	db.AutoMigrate(&Service{})
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
	service.Name = name
	service.Title = title
	service.Description = description
	service.Url = url
	service.Type = serviceType
	db.Save(&service)
}

type Reserved struct {
	gorm.Model
	AccountId      uint                  `gorm:"index:accountId_index"`
	Name           constants.ServiceName `json:"name"`
	Specifications common.PairList       `json:"specifications"`
}
