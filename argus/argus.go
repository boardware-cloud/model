package argus

import (
	"database/sql/driver"

	"github.com/boardware-cloud/common/constants"
	"github.com/boardware-cloud/common/utils"
	"gorm.io/gorm"
)

type Argus struct {
	gorm.Model
	AccountId   uint `gorm:"index:accountId_index"`
	Name        string
	Description string
	Status      constants.MonitorStatus
	Type        constants.MonitorType `gorm:"type:VARCHAR(128)"`
	ArgusNodeId *uint                 `gorm:"index:uptime_id_name"`
	Monitor     Monitor               `gorm:"type:JSON"`
}

func (a *Argus) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}

type Monitor interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

func Create(argus *Argus) {
	db.Save(argus)
}
