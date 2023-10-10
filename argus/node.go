package argus

import "gorm.io/gorm"

type ArgusNode struct {
	gorm.Model
	Heartbeat         int64
	HeartbeatInterval int64
}
