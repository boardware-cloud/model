package argus

import (
	"database/sql/driver"
	"encoding/json"
)

type PingMonitor struct {
	Interval int64  `json:"interval"`
	Timeout  int64  `json:"timeout"`
	Url      string `json:"url"`
	Retries  int64  `json:"retries"`
}

func (w *PingMonitor) Scan(value any) error {
	return json.Unmarshal(value.([]byte), w)
}

func (w PingMonitor) Value() (driver.Value, error) {
	b, err := json.Marshal(w)
	return b, err
}
