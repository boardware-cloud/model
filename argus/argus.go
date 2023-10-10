package argus

import (
	"database/sql/driver"

	"github.com/Dparty/common/utils"
	"github.com/boardware-cloud/common/constants"
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
