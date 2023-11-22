package argus

import (
	"github.com/boardware-cloud/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = model.GetDB()
	db.AutoMigrate(&Argus{}, &ArgusRecord{}, &ArgusNode{}, &NotificationRecord{})
}
