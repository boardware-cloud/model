package argus

import (
	"github.com/boardware-cloud/model/common"
	"github.com/boardware-cloud/model/core"
	"gorm.io/gorm"
)

var db *gorm.DB

var accountRepository core.AccountRepository

func Init(injectDB *gorm.DB) {
	db = injectDB
	db.AutoMigrate(&Argus{}, &ArgusRecord{}, &ArgusNode{}, &NotificationRecord{})
	accountRepository = core.NewAccountRepository(db)
	common.Init(injectDB)
}
