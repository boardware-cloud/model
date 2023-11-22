package common

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/boardware-cloud/common/code"
	"github.com/boardware-cloud/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = model.GetDB()
}

type PairList []Pair

func (PairList) GormDataType() string {
	return "JSON"
}

func (s *PairList) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s PairList) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type Pair struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

func (Pair) GormDataType() string {
	return "JSON"
}

func (s *Pair) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Pair) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type StringList []string

func (StringList) GormDataType() string {
	return "JSON"
}

func (s *StringList) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s StringList) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type List[T any] struct {
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Index int64 `json:"index"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}

func Find[T any](model T, conds ...any) (T, error) {
	if ctx := db.Find(model, conds...); ctx.RowsAffected == 0 {
		return model, code.ErrNotFound
	}
	return model, nil
}

func ListEntity(model any, index, limit int64, order string, where ...*gorm.DB) Pagination {
	ctx := db.Model(model)
	for _, w := range where {
		ctx = ctx.Where(w)
	}
	if order != "" {
		ctx = ctx.Order(order)
	}
	ctx = ctx.Limit(int(limit)).Offset(int(index * limit)).Find(model)
	var total int64
	ctx.Count(&total)
	return Pagination{
		Total: total,
		Index: index,
		Limit: limit,
	}
}
