package common

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init(injectDB *gorm.DB) {
	db = injectDB
}

type CurrentLimiting struct {
	gorm.Model
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

func ListModel[M any](model *[]M,
	index, limit int64,
	args ...any) List[M] {
	var total int64
	ctx := db.Model(model)
	if len(args) != 0 {
		ctx = ctx.Where(args[0], args[1:])
	}
	if total == 0 {
		return List[M]{
			Data: []M{},
			Pagination: Pagination{
				Limit: limit,
				Index: 0,
				Total: 0,
			},
		}
	}
	if total <= index*limit {
		index = total/limit - 1
	}
	ctx.Limit(int(limit)).Offset(int(index * limit)).Find(&model)
	return List[M]{
		Data: *model,
		Pagination: Pagination{
			Limit: limit,
			Index: index,
			Total: total,
		},
	}
}
