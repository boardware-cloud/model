package argus

import (
	"github.com/boardware-cloud/model/notification"
	"gorm.io/gorm"
)

type NotificationRecord struct {
	gorm.Model
	ArgusId           uint `gorm:"index:argus_id_index"`
	Message           string
	NotificationGroup notification.NotificationGroup
}
