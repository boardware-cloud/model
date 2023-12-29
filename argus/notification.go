package argus

import (
	"github.com/Dparty/common/singleton"
	"github.com/boardware-cloud/common/utils"
	"github.com/boardware-cloud/model"
	"github.com/boardware-cloud/model/notification"
	"gorm.io/gorm"
)

type NotificationRecord struct {
	gorm.Model
	ArgusId           uint `gorm:"index:argus_id_index"`
	Message           string
	NotificationGroup notification.NotificationGroup
}

func (a *NotificationRecord) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = utils.GenerteId()
	return err
}

var notificationRecordRepository = singleton.NewSingleton[NotificationRecordRepository](NewNotificationRecordRepository, singleton.Eager)

func GetNotificationRecordRepository() *NotificationRecordRepository {
	return notificationRecordRepository.Get()
}

func NewNotificationRecordRepository() *NotificationRecordRepository {
	return &NotificationRecordRepository{db: model.GetDB()}
}

type NotificationRecordRepository struct {
	db *gorm.DB
}
