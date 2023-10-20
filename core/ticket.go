package core

import (
	"github.com/boardware-cloud/common/utils"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	AccountId uint `json:"accountId" gorm:"index:account_id_index"`
	Type      string
	Secret    string
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = utils.GenerteId()
	return
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return TicketRepository{db: db}
}

type TicketRepository struct {
	db *gorm.DB
}
