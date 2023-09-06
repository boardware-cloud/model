package common

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

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
