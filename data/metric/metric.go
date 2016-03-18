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

func Init() Metrics {
	metrics := Metrics{}
	metrics.Items = []Metric{
		{"nodeping","status","/nodeping","get"},
		{"nan","status","/someOther","get"},
	}
	return metrics
}

func New(key string, typ string, endpoint string, method string) (*Metric) {
	return &Metric {
		key, typ, endpoint, method,
	}
}

func (m *Metric) Json() string {
	log.Println(m)
	a, err := json.Marshal(m)
	log.Println(err)
	log.Println(a)
	return string(a)
}

func (m *Metrics) Json() string {
	log.Println(m)
	a, err := json.Marshal(m)
	log.Println(err)
	log.Println(a)
	return string(a)
}
