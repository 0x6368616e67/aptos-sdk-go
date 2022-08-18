package client

import (
	"errors"
	"fmt"
)

const (
	ENInternal = "internal_error"
)

var (
	ErrNoResult = errors.New("no result inRPC response")
)

type ErrorMsg struct {
	Message            string `json:"message"`
	ErrorCode          string `json:"error_code"`
	AptosLedgerVersion string `json:"aptos_ledger_version"`
}

func (err ErrorMsg) Error() string {
	return fmt.Sprintf("%s[%s]: %s", err.ErrorCode, err.AptosLedgerVersion, err.Message)
}
