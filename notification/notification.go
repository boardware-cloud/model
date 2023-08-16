package notification

import (
	"github.com/boardware-cloud/common/utils"
	"github.com/boardware-cloud/model/common"
	"gorm.io/gorm"
)

type NotificationType string

type Email struct {
	gorm.Model
	Sender          string            `json:"sender"`
	Receivers       common.StringList `json:"receivers"`
	CarbonCopy      common.StringList `json:"carbonCopy"`
	BlindCarbonCopy common.StringList `json:"blindCarbonCopy"`
	Message         string            `json:"message"`
}

func (m *Email) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = utils.GenerteId()
	return
}
