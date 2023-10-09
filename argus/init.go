package argus

import (
	"github.com/boardware-cloud/model/common"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(injectDB *gorm.DB) {
	db = injectDB
	// db.AutoMigrate(&Monitor{})
	// db.AutoMigrate(&UptimeNode{})
	// db.AutoMigrate(&MonitoringRecord{})
	// db.AutoMigrate(&UptimeMonitorAlert{})
	// db.AutoMigrate(&ReservedMonitor{})
	db.AutoMigrate(&Argus{})
	common.Init(injectDB)
}
