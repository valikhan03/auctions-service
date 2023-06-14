package kafka

import(
	"encoding/json"
)

type Event struct {
	Command string                 `json:"command"`
	Entity  map[string]interface{} `json:"entity"`
}

func (e *Event) Marshal() (json.RawMessage, error) {
	return json.Marshal(e)
}
