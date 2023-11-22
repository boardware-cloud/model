package argus

import (
	"time"

	"github.com/boardware-cloud/common/utils"
	"github.com/boardware-cloud/model"
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

var argusRecordRepository *ArgusRecordRepository

func GetArgusRecordRepository() *ArgusRecordRepository {
	if argusRecordRepository == nil {
		argusRecordRepository = NewArgusRecordRepository()
	}
	return argusRecordRepository
}

func NewArgusRecordRepository() *ArgusRecordRepository {
	return &ArgusRecordRepository{model.GetDB()}
}

type ArgusRecordRepository struct {
	db *gorm.DB
}

func (a ArgusRecordRepository) Find(conds ...any) *ArgusRecord {
	var record ArgusRecord
	ctx := a.db.Find(&record, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &record
}

func (a ArgusRecordRepository) List(conds ...any) []ArgusRecord {
	var records []ArgusRecord
	a.db.Find(&records, conds...)
	return records
}
