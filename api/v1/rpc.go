package v1

type ErrorMsg struct {
	Message            string `json:"message"`
	ErrorCode          string `json:"error_code"`
	AptosLedgerVersion string `json:"aptos_ledger_version"`
}
