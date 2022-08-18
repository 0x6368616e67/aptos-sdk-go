package v1

const (
	BlockPath = "blocks/by_height/{block_height}"
)

type BlockReq struct {
	WithTransactions bool   `param:"with_transactions,omitempty"`
	BlockHeight      uint64 `path:"block_height"`
}

type BlockInfo struct {
	BlockHeight    string            `json:"block_height"`
	BlockHash      string            `json:"block_hash"`
	BlockTimestamp string            `json:"block_timestamp"`
	FirstVersion   string            `json:"first_version"`
	LastVersion    string            `json:"last_version"`
	Transactions   []TransactionInfo `json:"transactions"`
}
