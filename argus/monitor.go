package argus

import (
	"github.com/boardware-cloud/common/constants"
	"github.com/boardware-cloud/model/common"
	"gorm.io/gorm"
)

// import (
// 	"database/sql/driver"
// 	"encoding/json"
// 	"time"

// 	"github.com/boardware-cloud/common/constants"
// 	"github.com/boardware-cloud/common/utils"
// 	"github.com/boardware-cloud/model/common"
// 	"github.com/boardware-cloud/model/core"
// 	"gorm.io/gorm"
// )

// type UptimeNode struct {
// 	gorm.Model
// 	Heartbeat         int64
// 	HeartbeatInterval int64
// }

// func (m *UptimeNode) BeforeCreate(tx *gorm.DB) (err error) {
// 	m.ID = utils.GenerteId()
// 	return
// }

// type ReservedMonitor struct {
// 	gorm.Model
// 	AccountId uint `gorm:"index:accountId_index"`
// 	StartAt   time.Time
// 	ExpiredAt time.Time
// }

type Monitor2 struct {
	gorm.Model
	AccountId     uint `gorm:"index:accountId_index"`
	Name          string
	Description   string
	Status        constants.MonitorStatus
	Interval      int64
	Timeout       int64
	BaseTime      int64
	Url           string
	Retries       int64
	Type          constants.MonitorType `gorm:"type:VARCHAR(128)"`
	HttpMethod    *constants.HttpMehotd `gorm:"type:VARCHAR(128)"`
	UptimeNodeId  *uint                 `gorm:"index:uptime_id_name"`
	Heartbeat     int64
	AlertInterval int64
	// Notifications        Notifications `json:"notifications"`
	NotificationInterval int64 `json:"notificationInterval"`
	// Records              []MonitoringRecord
	// Alert                []UptimeMonitorAlert
	BodyRaw             *string                 `json:"bodyRaw"`
	BodyForm            *constants.HttpBodyForm `json:"bodyForm"`
	ContentType         *constants.ContentType  `json:"contentType"`
	Headers             *common.PairList        `json:"headers"`
	AcceptedStatusCodes *common.StringList      `json:"acceptedStatusCodes"`
}

// func (m Monitor) Owner() core.Account {
// 	return Owner(m)
// }

// func (m *Monitor) Off() {
// 	Off(m)
// }

// func Owner(m Monitor) core.Account {
// 	owner, _ := core.FindAccount(m.AccountId)
// 	return owner
// }

// func Off(m *Monitor) {
// 	m.Status = constants.DISACTIVED
// 	db.Save(m)
// }

// func (m *Monitor) BeforeCreate(tx *gorm.DB) (err error) {
// 	m.ID = utils.GenerteId()
// 	return
// }

// type MonitoringRecord struct {
// 	gorm.Model
// 	MonitorId    uint `gorm:"index:monitorId_index"`
// 	CheckedAt    time.Time
// 	Url          string
// 	Type         constants.MonitorType `gorm:"type:VARCHAR(128)"`
// 	HttpMethod   *constants.HttpMehotd `gorm:"type:VARCHAR(128)"`
// 	StatusCode   string
// 	Result       constants.MonitoringResult `gorm:"type:VARCHAR(128)"`
// 	ResponseTime *int64
// 	Body         *string                 `json:"body"`
// 	Headers      *common.PairList        `json:"headers"`
// 	BodyForm     *constants.HttpBodyForm `json:"bodyForm"`
// 	ContentType  *constants.ContentType  `json:"contentType"`
// }

// func (m *MonitoringRecord) BeforeCreate(tx *gorm.DB) (err error) {
// 	m.ID = utils.GenerteId()
// 	return
// }

// type EmailReceivers struct {
// 	gorm.Model
// 	To  common.StringList `json:"to"`
// 	Cc  common.StringList `json:"cc"`
// 	Bcc common.StringList `json:"bcc"`
// }

// func (m *EmailReceivers) BeforeCreate(tx *gorm.DB) (err error) {
// 	m.ID = utils.GenerteId()
// 	return
// }

// func (s *EmailReceivers) Scan(value any) error {
// 	return json.Unmarshal(value.([]byte), s)
// }

// func (s EmailReceivers) Value() (driver.Value, error) {
// 	b, err := json.Marshal(s)
// 	return b, err
// }

// func (EmailReceivers) GormDataType() string {
// 	return "JSON"
// }

// type Notifications []Notification

// func (Notifications) GormDataType() string {
// 	return "JSON"
// }

// func (s *Notifications) Scan(value any) error {
// 	return json.Unmarshal(value.([]byte), s)
// }

// func (s Notifications) Value() (driver.Value, error) {
// 	b, err := json.Marshal(s)
// 	return b, err
// }

// type Notification struct {
// 	gorm.Model
// 	Type           constants.NotificationType `json:"type" gorm:"type:VARCHAR(128)"`
// 	EmailReceivers *EmailReceivers            `json:"emailReceivers"`
// }

// func (m *Notification) BeforeCreate(tx *gorm.DB) (err error) {
// 	m.ID = utils.GenerteId()
// 	return
// }

// func (s *Notification) Scan(value any) error {
// 	return json.Unmarshal(value.([]byte), s)
// }

// func (s Notification) Value() (driver.Value, error) {
// 	b, err := json.Marshal(s)
// 	return b, err
// }

// func (Notification) GormDataType() string {
// 	return "JSON"
// }

// type Email struct {
// 	EmailReceivers
// 	Message string `json:"message"`
// }

// func (Email) GormDataType() string {
// 	return "JSON"
// }

// func (s *Email) Scan(value any) error {
// 	return json.Unmarshal(value.([]byte), s)
// }

// func (s Email) Value() (driver.Value, error) {
// 	b, err := json.Marshal(s)
// 	return b, err
// }

// type UptimeMonitorAlert struct {
// 	gorm.Model
// 	MonitorId     uint          `gorm:"index:monitor_id_index"`
// 	Notifications Notifications `json:"email"`
// 	Subject       string        `json:"subject"`
// 	Message       string        `json:"message"`
// }

// func (m *UptimeMonitorAlert) BeforeCreate(tx *gorm.DB) (err error) {
// 	m.ID = utils.GenerteId()
// 	return
// }
