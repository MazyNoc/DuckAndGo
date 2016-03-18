package metric

import (
	"log"
	"encoding/json"
)

type Metric struct {
	Key string `json:"key"`
	Typ string `json:"type"`
	Endpoint string `json:"endpoint"`
	Method string `json:"method"`
}

type Metrics struct {
	Items []Metric `json:"metrics"`
}

func Get(host string) Metrics {
	metrics := Metrics{}
	metrics.Items = []Metric{
		{"nodeping","status",host + "/nodeping","get"},
		{"Not used","status",host + "/someOther","get"},
	}
	return metrics
}

func (m *Metric) Json() string {
	log.Println(m)
	result, _ := json.Marshal(m)
	return string(result)
}

func (m *Metrics) Json() string {
	log.Println(m)
	result, _ := json.Marshal(m)
	return string(result)
}
