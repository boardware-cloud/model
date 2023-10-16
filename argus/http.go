package argus

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/boardware-cloud/common/constants"
	"github.com/boardware-cloud/model/common"
)

type HttpMonitor struct {
	Type                string               `json:"type"`
	Interval            time.Duration        `json:"interval"`
	Timeout             int64                `json:"timeout"`
	Url                 string               `json:"url"`
	Retries             int64                `json:"retries"`
	HttpMethod          constants.HttpMehotd `json:"method" gorm:"type:VARCHAR(128)"`
	Headers             common.PairList      `json:"headers"`
	AcceptedStatusCodes common.StringList    `json:"acceptedStatusCodes"`
}

func (h HttpMonitor) Target() string {
	return h.Url
}

func (h HttpMonitor) ToJSON() MonitorJSON {
	b, _ := json.Marshal(h)
	return b
}

func (h *HttpMonitor) Scan(value any) error {
	return json.Unmarshal(value.([]byte), h)
}

func (h HttpMonitor) Value() (driver.Value, error) {
	b, err := json.Marshal(h)
	return b, err
}

func (h HttpMonitor) GetType() string {
	return "HTTP"
}
