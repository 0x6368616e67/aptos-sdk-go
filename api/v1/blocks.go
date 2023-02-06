package v1

const (
	BlockByHeightPath  = "GET@blocks/by_height/{height}"
	BlockByVersionPath = "GET@blocks/by_version/{version}"
)

type BlockByHeightReq struct {
	WithTransactions bool   `param:"with_transactions,omitempty"`
	Height           uint64 `path:"height"`
}

type BlockByVersionReq struct {
	WithTransactions bool   `param:"with_transactions,omitempty"`
	Version          uint64 `path:"version"`
}

type BlockInfo struct {
	BlockHeight    string            `json:"block_height"`
	BlockHash      string            `json:"block_hash"`
	BlockTimestamp string            `json:"block_timestamp"`
	FirstVersion   string            `json:"first_version"`
	LastVersion    string            `json:"last_version"`
	Transactions   []TransactionInfo `json:"transactions"`
}
