package client

import (
	"encoding/json"
	"fmt"
)

type rpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err *rpcError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("rpc error %d", err.Code)
	}
	return err.Message
}

func (err *rpcError) ErrorCode() int {
	return err.Code
}

func (err *rpcError) ErrorData() interface{} {
	return err.Data
}

type rpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}
