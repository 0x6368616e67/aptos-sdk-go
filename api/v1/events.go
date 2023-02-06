package v1

import "encoding/json"

const (
	EventByCreationNumberPath = "GET@accounts/{address}/events/{creation_number}"
	EventByEventHandlerPath   = "GET@accounts/{address}/events/{event_handle}/{field_name}"
)

type EventByCreationNumberReq struct {
	Limit          uint16 `param:"limit,omitempty"`
	Start          uint64 `param:"start,omitempty"`
	Address        string `path:"address"`
	CreationNumber string `path:"creation_number"`
}

type EventByEventHandlerReq struct {
	Address string `path:"address"`
	Handler string `path:"event_handle"`
	Filed   string `path:"field_name"`

	Limit uint16 `param:"limit,omitempty"`
	Start uint64 `param:"start,omitempty"`
}

type GUIDInfo struct {
	AccountAddress string `path:"account_address"`
	CreationNumber string `path:"creation_number"`
}

type EventInfo struct {
	Version        string          `json:"version"`
	GUID           GUIDInfo        `json:"guid"`
	SequenceNumber string          `json:"sequence_number"`
	Type           string          `json:"type"`
	Data           json.RawMessage `json:"data"`
}
