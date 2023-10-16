package argus

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/boardware-cloud/common/constants"
	"github.com/boardware-cloud/common/utils"
	"github.com/boardware-cloud/model/abstract"
	"github.com/boardware-cloud/model/common"
	"github.com/boardware-cloud/model/core"
	"github.com/boardware-cloud/model/notification"
	"gorm.io/gorm"
)

type Argus struct {
	gorm.Model
	AccountId         uint `gorm:"index:accountId_index"`
	Name              string
	Description       string
	Status            constants.MonitorStatus
	Type              constants.MonitorType `gorm:"type:VARCHAR(128)"`
	ArgusNodeId       *uint                 `gorm:"index:uptime_id_name"`
	MonitorJSON       MonitorJSON           `gorm:"type:JSON"`
	NotificationGroup notification.NotificationGroup
}

func (a Argus) LastNotificationRecord() *NotificationRecord {
	record, err := common.Find(&NotificationRecord{})
	if err != nil {
		return nil
	}
	return record
}

func (a *Argus) Update(n Argus) {
	a.UpdatedAt = time.Now()
	a.Name = n.Name
	a.Description = n.Description
	a.Status = n.Status
	a.Type = n.Type
	a.MonitorJSON = n.MonitorJSON
	db.Save(a)
}

func (a *Argus) Spawn(nodeId uint) bool {
	a.ArgusNodeId = &nodeId
	ctx := db.Save(a)
	return ctx.RowsAffected != 0
}

func (a Argus) Record(result string, responesTime time.Duration) ArgusRecord {
	record := ArgusRecord{Result: result, ArgusId: a.ID, ResponesTime: responesTime}
	db.Save(&record)
	return record
}

func (a Argus) LastRecord() *ArgusRecord {
	record, err := common.Find(&ArgusRecord{}, "argus_id = ?", a.ID)
	if err != nil {
		return nil
	}
	return record
}

func (a Argus) Owner() abstract.Owner {
	var account core.Account
	db.Find(&account, a.AccountId)
	return account
}

func (a Argus) Monitor() Monitor {
	return a.MonitorJSON.Monitor()
}

func (a *Argus) SetMonitor(monitor Monitor) Argus {
	a.MonitorJSON = monitor.ToJSON()
	return *a
}

func (a *Argus) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}

type Monitor interface {
	GetType() string
	ToJSON() MonitorJSON
	Scan(value any) error
	Value() (driver.Value, error)
}

type MonitorJSON json.RawMessage

func (j *MonitorJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}
	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = MonitorJSON(result)
	return err
}

func (j MonitorJSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

func (m MonitorJSON) Monitor() Monitor {
	ma := make(map[string]interface{})
	json.Unmarshal(m, &ma)
	switch ma["type"] {
	case "HTTP":
		var httpMonitor HttpMonitor
		json.Unmarshal(m, &httpMonitor)
		return &httpMonitor
	case "PING":
		var pingMonitor PingMonitor
		json.Unmarshal(m, &pingMonitor)
		return &pingMonitor
	}
	return nil
}

func NewArgusRepository(db *gorm.DB) ArgusRepository {
	return ArgusRepository{db}
}

type ArgusRepository struct {
	db *gorm.DB
}

func (a ArgusRepository) Find(conds ...any) *Argus {
	var argus Argus
	ctx := a.db.Find(&argus, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &argus
}

func (a ArgusRepository) GetById(id uint) *Argus {
	return a.Find(id)
}
