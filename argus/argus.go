package argus

import (
	"database/sql/driver"
	"encoding/json"

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
	Monitor     MonitorEnitiy         `gorm:"type:JSON"`
}

func (a *Argus) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}

type MonitorEnitiy struct {
	Type    constants.MonitorType `json:"type"`
	Monitor Monitor               `json:"monitor"`
}

func (MonitorEnitiy) GormDataType() string {
	return "JSON"
}

func (w *MonitorEnitiy) Scan(value any) error {
	m := make(map[string]any)
	json.Unmarshal(value.([]byte), &m)
	w.Type = m["type"].(constants.MonitorType)
	switch w.Type {
	case constants.HTTP:
		monitor := m["monitor"].(HttpMonitor)
		w.Monitor = &monitor
	case constants.PING:
		monitor := m["monitor"].(PingMonitor)
		w.Monitor = &monitor
	}
	return nil
}

func (w MonitorEnitiy) Value() (driver.Value, error) {
	b, err := json.Marshal(w)
	return b, err
}

type Monitor interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

func Create(argus *Argus) {
	db.Save(argus)
}
