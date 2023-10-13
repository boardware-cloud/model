package argus

import (
	"github.com/boardware-cloud/model/common"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(injectDB *gorm.DB) {
	db = injectDB
	db.AutoMigrate(&Argus{}, &ArgusRecord{}, &ArgusNode{}, &NotificationRecord{})
	common.Init(injectDB)
}
