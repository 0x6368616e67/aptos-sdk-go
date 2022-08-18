package v1

const (
	AccountPath = "accounts/{address}"
)

type AccountReq struct {
	LedgerVersion uint64 `param:"ledger_version,omitempty" `
	Address       string `path:"address"`
}

type AccountInfo struct {
	SequenceNumber    string `json:"sequence_number"`
	AuthenticationKey string `json:"authentication_key"`
}
