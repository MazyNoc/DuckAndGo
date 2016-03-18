package sample
// single datum type

import (
	"time"
	"encoding/json"
)

type Datum struct {
	Status string `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type Sample struct {
	Key string `json:"key"`
	Typ string `json:"type"`
	Data Datum `json:"data"`
}

func Init() Sample {
	sample := Sample{"sample","testing", Datum{"ok", time.Now()}}
	return sample
}

func (m *Sample) Json() string {
	result, _ := json.Marshal(m)
	return string(result)
}

func ConvertStatus(value int) string{
	switch value {
	case 1:
		return "ok"
	default:
		return "error"

	}
}