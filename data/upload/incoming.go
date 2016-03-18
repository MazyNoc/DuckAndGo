package upload

import (
	"encoding/json"
)
type Nodeping struct {
	Value int `json:"value"`
}

func (m *Nodeping) Json() string {
	a, _ := json.Marshal(m)
	return string(a)
}
