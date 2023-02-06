package v1

import "encoding/json"

const (
	TransactionPath          = "GET@transactions"
	TransactionByHashPath    = "GET@transactions/by_hash/{txn_hash}"
	TransactionByVersionPath = "GET@transactions/by_version/{txn_version}"
	TransactionOfAccountPath = "GET@accounts/{address}/transactions"
	TransactionEncodingPath  = "POST@transactions/encode_submission"
	TransactionSimulatePath  = "POST@transactions/simulate"
	TransactionSubmitPath    = "POST@transactions"
	EstimateGasPricePath     = "GET@estimate_gas_price"
)

type EstimateGasPrice struct {
	DeprioritizedGasEstimate int `json:"deprioritized_gas_estimate"`
	GasEstimate              int `json:"gas_estimate"`
	PrioritizedGasEstimate   int `json:"prioritized_gas_estimate"`
}

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
	Address      string                    `json:"address"`
	StateKeyHash string                    `json:"state_key_hash"`
	Type         string                    `json:"type"`
	Data         TransactionChangeDataInfo `json:"data"`
}

type TransactionPayload struct {
	Function      string            `json:"function"`
	TypeArguments []string          `json:"type_arguments"`
	Arguments     []json.RawMessage `json:"arguments"`
	Type          string            `json:"type"`
}

type TransactionSignature struct {
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
	Type      string `json:"type"`
}

type TransactionInfo struct {
	Version                  string               `json:"version"`
	Hash                     string               `json:"hash"`
	StateChangeHash          string               `json:"state_change_hash"`
	EventRootHash            string               `json:"event_root_hash"`
	StateCheckpointHash      string               `json:"state_checkpoint_hash"`
	GasUsed                  string               `json:"gas_used"`
	Success                  bool                 `json:"success"`
	VMStatus                 string               `json:"vm_status"`
	AccumulatorRootHash      string               `json:"accumulator_root_hash"`
	Changes                  []json.RawMessage    `json:"changes"`
	ID                       string               `json:"id,omitempty"`
	Epoch                    string               `json:"epoch,omitempty"`
	Round                    string               `json:"round,omitempty"`
	Events                   []EventInfo          `json:"events,omitempty"`
	PreviousBlockVotesBitvec []bool               `json:"previous_block_votes_bitvec,omitempty"`
	Proposer                 string               `json:"proposer,omitempty"`
	FailedProposerIndices    []json.RawMessage    `json:"failed_proposer_indices,omitempty"`
	Timestamp                string               `json:"timestamp"`
	Type                     string               `json:"type"`
	Sender                   string               `json:"sender,omitempty"`
	SequenceNumber           string               `json:"sequence_number,omitempty"`
	MaxGasAmount             string               `json:"max_gas_amount,omitempty"`
	GasUnitPrice             string               `json:"gas_unit_price,omitempty"`
	ExpirationTimestampSecs  string               `json:"expiration_timestamp_secs,omitempty"`
	Payload                  TransactionPayload   `json:"payload,omitempty"`
	Signature                TransactionSignature `json:"signature,omitempty"`
}
