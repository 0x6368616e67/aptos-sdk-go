package v1

const (
	BlockPath = "blocks/by_height/{block_height}"
)

type BlockReq struct {
	WithTransactions bool   `param:"with_transactions,omitempty"`
	BlockHeight      uint64 `path:"block_height"`
}

type TransactionGUID struct {
	ID struct {
		Addr        string `json:"addr"`
		CreationNum string `json:"creation_num"`
	} `json:"id"`
}

type TransactionChangeDataEvent struct {
	Counter string          `json:"counter"`
	GUID    TransactionGUID `json:"guid"`
}

type TransactionChangeDataInfo struct {
	EpochInterval  string                     `json:"epoch_interval"`
	Height         string                     `json:"height"`
	NewBlockEvents TransactionChangeDataEvent `json:"new_block_events"`
}

type TransactionChangeData struct {
	Type string                    `json:"type"`
	Data TransactionChangeDataInfo `json:"data"`
}

type TransactionChangeInfo struct {
	Address      string                `json:"address"`
	StateKeyHash string                `json:"state_key_hash"`
	Data         TransactionChangeData `json:"data"`
	Type         string                `json:"type"`
}

type TransactionEventData struct {
	Epoch                 string        `json:"epoch"`
	FailedProposerIndices []interface{} `json:"failed_proposer_indices"`
	Height                string        `json:"height"`
	PreviousBlockVotes    []bool        `json:"previous_block_votes"`
	Proposer              string        `json:"proposer"`
	Round                 string        `json:"round"`
	TimeMicroseconds      string        `json:"time_microseconds"`
}

type TransactionEventInfo struct {
	Key            string               `json:"key"`
	SequenceNumber string               `json:"sequence_number"`
	Type           string               `json:"type"`
	Data           TransactionEventData `json:"data"`
}

type TransactionInfo struct {
	Version               string                  `json:"version"`
	Hash                  string                  `json:"hash"`
	StateRootHash         string                  `json:"state_root_hash"`
	EventRootHash         string                  `json:"event_root_hash"`
	GasUsed               string                  `json:"gas_used"`
	Success               bool                    `json:"success"`
	VMStatus              string                  `json:"vm_status"`
	AccumulatorRootHash   string                  `json:"accumulator_root_hash"`
	Changes               []TransactionChangeInfo `json:"changes"`
	ID                    string                  `json:"id,omitempty"`
	Epoch                 string                  `json:"epoch,omitempty"`
	Round                 string                  `json:"round,omitempty"`
	Events                []TransactionEventInfo  `json:"events,omitempty"`
	PreviousBlockVotes    []bool                  `json:"previous_block_votes,omitempty"`
	Proposer              string                  `json:"proposer,omitempty"`
	FailedProposerIndices []interface{}           `json:"failed_proposer_indices,omitempty"`
	Timestamp             string                  `json:"timestamp"`
	Type                  string                  `json:"type"`
}

type BlockInfo struct {
	BlockHeight    string            `json:"block_height"`
	BlockHash      string            `json:"block_hash"`
	BlockTimestamp string            `json:"block_timestamp"`
	FirstVersion   string            `json:"first_version"`
	LastVersion    string            `json:"last_version"`
	Transactions   []TransactionInfo `json:"transactions"`
}
