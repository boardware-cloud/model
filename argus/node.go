package argus

import (
	"time"

	"github.com/boardware-cloud/common/utils"
	"gorm.io/gorm"
)

type ArgusNode struct {
	gorm.Model
	Heartbeat         int64
	HeartbeatInterval time.Duration
}

func (a *ArgusNode) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return
}
