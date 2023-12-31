package argus

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type PingMonitor struct {
	Type     string        `json:"type"`
	Interval time.Duration `json:"interval"`
	Host     string        `json:"host"`
	Retries  int64         `json:"retries"`
}

func (p PingMonitor) Target() string {
	return p.Host
}

func (h PingMonitor) ToJSON() MonitorJSON {
	b, _ := json.Marshal(h)
	return b
}

func (w *PingMonitor) Scan(value any) error {
	return json.Unmarshal(value.([]byte), w)
}

func (w PingMonitor) Value() (driver.Value, error) {
	b, err := json.Marshal(w)
	return b, err
}

func (h PingMonitor) GetType() string {
	return "PING"
}
