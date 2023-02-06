package v1

const (
	HealthyPath = "GET@-/healthy"
	LedgerPath  = "GET@"
)

type HealthyReq struct {
	Duration uint32 `param:"duration_secs,omitempty"`
}

type HealthyRsp struct {
	Message string `json:"message"`
}

type LedgerInfo struct {
	ChainID             int    `json:"chain_id"`
	Epoch               string `json:"epoch"`
	LedgerVersion       string `json:"ledger_version"`
	OldestLedgerVersion string `json:"oldest_ledger_version"`
	LedgerTimestamp     string `json:"ledger_timestamp"`
	NodeRole            string `json:"node_role"`
	OldestBlockHeight   string `json:"oldest_block_height"`
	BlockHeight         string `json:"block_height"`
	GitHash             string `json:"git_hash"`
}

type LedgerRsp struct {
	LedgerInfo
}
