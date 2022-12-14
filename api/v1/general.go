package v1

const (
	HealthyPath = "GET@-/healthy"
	LedgerPath  = "GET@"
)

type HealthyReq struct {
	Duration uint32 `param:"duration_secs,omitempty"`
}

type HealthyRsp struct {
}

type LedgerInfo struct {
	ChainID             int    `json:"chain_id"`
	Epoch               string `json:"epoch"`
	LedgerVersion       string `json:"ledger_version"`
	OldestLedgerVersion string `json:"oldest_ledger_version"`
	BlockHeight         string `json:"block_height"`
	OldestBlockHeight   string `json:"oldest_block_height"`
	LedgerTimestamp     string `json:"ledger_timestamp"`
	NodeRole            string `json:"node_role"`
}

type LedgerRsp struct {
	LedgerInfo
}
