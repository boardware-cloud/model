package argus

import (
	"time"

	"github.com/boardware-cloud/common/utils"
	"gorm.io/gorm"
)

type ArgusRecord struct {
	gorm.Model
	ArgusId      uint `gorm:"index:argus_index"`
	Result       string
	ResponesTime time.Duration
}

func (a *ArgusRecord) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}
