package billing

import (
	"github.com/boardware-cloud/model/core"
	"gorm.io/gorm"
)

type Quantifiable interface {
	Owner() core.Account
	Pricing() int64
	Off()
	Service() Service
	Reserved() bool
	MetaData() string
}

type Subtotal struct {
	gorm.Model
	AccountId uint `gorm:"index:accountId_index"`
}
