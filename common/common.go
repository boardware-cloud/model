package common

import (
	"database/sql/driver"
	"encoding/json"
)

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
