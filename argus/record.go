package argus

import "gorm.io/gorm"

type ArgusRecord struct {
	gorm.Model
	ArgusId uint `gorm:"index:argus_index"`
	Result  string
}
