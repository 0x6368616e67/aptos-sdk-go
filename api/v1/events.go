package v1

import "encoding/json"

const (
	EventPath        = "events/{event_key}"
	EventWithHandler = "accounts/{address}/events/{event_handle}/{field_name}"
)

type EventReq struct {
	Limit    uint16 `param:"limit,omitempty"`
	Start    uint64 `param:"start,omitempty"`
	EventKey string `path:"event_key"`
}

type EventWithHandlerReq struct {
	Limit   uint16 `param:"limit,omitempty"`
	Address string `path:"address"`
	Handler string `path:"event_handle"`
	Filed   string `path:"field_name"`
}

type EventInfo struct {
	Version        string          `json:"version"`
	Key            string          `json:"key"`
	SequenceNumber string          `json:"sequence_number"`
	Type           string          `json:"type"`
	Data           json.RawMessage `json:"data"`
}
