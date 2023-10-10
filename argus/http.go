package argus

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/boardware-cloud/common/constants"
)

type HttpMonitor struct {
	Interval   int64                `json:"interval"`
	Timeout    int64                `json:"timeout"`
	Url        string               `json:"url"`
	Retries    int64                `json:"retries"`
	HttpMethod constants.HttpMehotd `json:"method" gorm:"type:VARCHAR(128)"`
}

func (w *HttpMonitor) Scan(value any) error {
	return json.Unmarshal(value.([]byte), w)
}

func (w HttpMonitor) Value() (driver.Value, error) {
	b, err := json.Marshal(w)
	return b, err
}
