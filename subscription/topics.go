package subscription

import "encoding/json"

type Topic interface {
	Stringify() ([]byte, error)
}

type NewBlockQuery struct {
	Query string `json:"query"`
}

func (query *NewBlockQuery) Stringify() ([]byte, error) {
	return json.Marshal(query)
}

type NewEventLogQuery struct {
	Query     string   `json:"query"`
	Addresses []string `json:"addresses"`
}

func (query *NewEventLogQuery) Stringify() ([]byte, error) {
	return json.Marshal(query)
}

type Unsubscribe struct {
	Query string `json:"query"`
	Type  string `json:"type"`
}

func (query *Unsubscribe) Stringify() ([]byte, error) {
	return json.Marshal(query)
}
