package core

import (
	"math/rand"
	"strings"
	"time"

	errorCode "github.com/boardware-cloud/common/code"
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
	var ticket Ticket = Ticket{
		AccountId: accountId,
		Secret:    RandomNumberString(16),
		Type:      typ,
	}
	t.db.Save(&ticket)
	return ticket
}

func (t TicketRepository) UseTicket(token string) (Ticket, error) {
	ss := strings.Split(token, ":")
	var ticket Ticket
	if len(ss) != 2 {
		return ticket, errorCode.ErrUnauthorized
	}
	ctx := t.db.Find(&ticket, utils.StringToUint(ss[0]))
	if ctx.RowsAffected == 0 {
		return ticket, errorCode.ErrUnauthorized
	}
	if ticket.Secret != ss[1] {
		return ticket, errorCode.ErrUnauthorized
	}
	t.db.Delete(&ticket)
	if time.Now().Unix()-ticket.CreatedAt.Unix() > 300 {
		return ticket, errorCode.ErrUnauthorized
	}
	return ticket, nil
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
