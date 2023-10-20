package core

import (
	"math/rand"
	"time"

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

func (t TicketRepository) CreateTicket(typ string, accountId uint) Ticket {
	var ticket Ticket
	ticket.AccountId = accountId
	ticket.Secret = RandomNumberString(16)
	ticket.Type = typ
	t.db.Save(&ticket)
	return ticket
}

const charset = "0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomNumberString(length int) string {
	return StringWithCharset(length, charset)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
