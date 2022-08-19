package v1

const (
	TransactionPath          = "GET@transactions"
	TransactionByHashPath    = "GET@transactions/by_hash/{txn_hash}"
	TransactionByVersionPath = "GET@transactions/by_version/{txn_version}"
	TransactionOfAccountPath = "GET@accounts/{address}/transactions"
	TransactionEncodingPath  = "POST@transactions/encode_submission"
	TransactionSimulatePath  = "POST@transactions/simulate"
	TransactionSubmitPath    = "POST@transactions"
)

type TransactionOfAccountReq struct {
	Address string `path:"address"`
	Limit   uint16 `param:"limit,omitempty"`
	Start   uint64 `param:"start,omitempty"`
}

type TransactionByVersionReq struct {
	Version uint64 `path:"txn_version"`
}

type TransactionByHashReq struct {
	Hash string `path:"txn_hash"`
}

type TransactionReq struct {
	Limit uint16 `param:"limit,omitempty"`
	Start uint64 `param:"start,omitempty"`
}

type TransactionID struct {
	Addr        string `json:"addr"`
	CreationNum string `json:"creation_num"`
}

type TransactionGUID struct {
	ID TransactionID `json:"id"`
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

type TransactionChangeInfo struct {
	Address      string                `json:"address"`
	StateKeyHash string                `json:"state_key_hash"`
	Data         TransactionChangeData `json:"data"`
	Type         string                `json:"type"`
}

type TransactionChangeData struct {
	Type string                    `json:"type"`
	Data TransactionChangeDataInfo `json:"data"`
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

type TransactionPayload struct {
	Function      string   `json:"function"`
	TypeArguments []string `json:"type_arguments"`
	Arguments     []string `json:"arguments"`
	Type          string   `json:"type"`
}

type TransactionSignature struct {
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
	Type      string `json:"type"`
}

type TransactionInfo struct {
	Version                 string                 `json:"version"`
	Hash                    string                 `json:"hash"`
	StateRootHash           string                 `json:"state_root_hash"`
	EventRootHash           string                 `json:"event_root_hash"`
	GasUsed                 string                 `json:"gas_used"`
	Success                 bool                   `json:"success"`
	VMStatus                string                 `json:"vm_status"`
	AccumulatorRootHash     string                 `json:"accumulator_root_hash"`
	Changes                 TransactionChangeData  `json:"changes"`
	ID                      string                 `json:"id,omitempty"`
	Epoch                   string                 `json:"epoch,omitempty"`
	Round                   string                 `json:"round,omitempty"`
	Events                  []TransactionEventInfo `json:"events,omitempty"`
	PreviousBlockVotes      []bool                 `json:"previous_block_votes,omitempty"`
	Proposer                string                 `json:"proposer,omitempty"`
	FailedProposerIndices   []interface{}          `json:"failed_proposer_indices,omitempty"`
	Timestamp               string                 `json:"timestamp"`
	Type                    string                 `json:"type"`
	Sender                  string                 `json:"sender,omitempty"`
	SequenceNumber          string                 `json:"sequence_number,omitempty"`
	MaxGasAmount            string                 `json:"max_gas_amount,omitempty"`
	GasUnitPrice            string                 `json:"gas_unit_price,omitempty"`
	ExpirationTimestampSecs string                 `json:"expiration_timestamp_secs,omitempty"`
	Payload                 TransactionPayload     `json:"payload,omitempty"`
	Signature               TransactionSignature   `json:"signature,omitempty"`
}
