package notification

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// var db *gorm.DB

// func Init(injectDB *gorm.DB) {
// 	db = injectDB
// 	db.AutoMigrate()
// }

type NotificationGroup struct {
	Interval          time.Duration   `json:"interval"`
	NotificationsJSON json.RawMessage `json:"JSON"`
}

func (NotificationGroup) GormDataType() string {
	return "JSON"
}

func (h *NotificationGroup) Scan(value any) error {
	return json.Unmarshal(value.([]byte), h)
}

func (h NotificationGroup) Value() (driver.Value, error) {
	b, err := json.Marshal(h)
	return b, err
}

func (g *NotificationGroup) SetNotifications(n []Notification) *NotificationGroup {
	j, _ := json.Marshal(n)
	g.NotificationsJSON = j
	return g
}

func (g NotificationGroup) Notifications() []Notification {
	var notifications []Notification
	json.Unmarshal(g.NotificationsJSON, &notifications)
	return notifications
}

type Notification struct {
	Interval   *time.Duration
	Type       string
	EntityJSON json.RawMessage `json:"JSON"`
}

func (n *Notification) SetEntity(entity Entity) *Notification {
	n.Type = entity.Type()
	j, _ := json.Marshal(entity)
	n.EntityJSON = j
	return n
}

func (n Notification) Entity() Entity {
	switch n.Type {
	case "EMAIL":
		var email Email
		json.Unmarshal(n.EntityJSON, &email)
		return email
	}
	return nil
}

type Entity interface {
	Type() string
}

type Email struct {
	gorm.Model
	To  []string `json:"to"`
	Cc  []string `json:"cc"`
	Bcc []string `json:"bcc"`
}

func (Email) Type() string {
	return "Email"
}
