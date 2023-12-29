package argus

import (
	"time"

	"github.com/Dparty/common/singleton"
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

var argusRecordRepository = singleton.NewSingleton[ArgusRecordRepository](NewArgusRecordRepository, singleton.Eager)

func GetArgusRecordRepository() *ArgusRecordRepository {
	return argusRecordRepository.Get()
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
