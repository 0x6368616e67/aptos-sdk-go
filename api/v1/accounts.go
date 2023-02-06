package v1

import "encoding/json"

const (
	AccountPath                 = "GET@accounts/{address}"
	AccountResourcePath         = "GET@accounts/{address}/resources"
	AccountModulePath           = "GET@accounts/{address}/modules"
	AccountModuleWithNamePath   = "GET@accounts/{address}/module/{module_name}"
	AccountResourceWithTypePath = "GET@accounts/{address}/resource/{resource_type}"
)

type AccountReq struct {
	LedgerVersion uint64 `param:"ledger_version,omitempty" `
	Address       string `path:"address"`
}

type AccountInfo struct {
	SequenceNumber    string `json:"sequence_number"`
	AuthenticationKey string `json:"authentication_key"`
}

type AccountResourceReq struct {
	LedgerVersion uint64 `param:"ledger_version,omitempty" `
	Address       string `path:"address"`
	Start         int64  `param:"start,omitempty" `
	Limit         int64  `param:"limit,omitempty" `
}

type AccountResourceWithTypeReq struct {
	LedgerVersion uint64 `param:"ledger_version,omitempty" `
	Address       string `path:"address"`
	Type          string `path:"resource_type"`
}

type AccountResourceInfo struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type AccountModuleReq struct {
	LedgerVersion uint64 `param:"ledger_version,omitempty" `
	Address       string `path:"address"`
	Start         int64  `param:"start,omitempty" `
	Limit         int64  `param:"limit,omitempty" `
}

type AccountModuleWithNameReq struct {
	LedgerVersion uint64 `param:"ledger_version,omitempty" `
	Address       string `path:"address"`
	Name          string `path:"module_name"`
}

type GenericTypeParam struct {
	Constraints []string `json:"constraints"`
}

type ExposedFunction struct {
	Name              string             `json:"name"`
	Visibility        string             `json:"visibility"`
	IsEntry           bool               `json:"is_entry"`
	GenericTypeParams []GenericTypeParam `json:"generic_type_params"`
	Params            []string           `json:"params"`
	Return            []string           `json:"return"`
}

type ModuleStructField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ModuleStruct struct {
	Name              string              `json:"name"`
	IsNative          bool                `json:"is_native"`
	Abilities         []string            `json:"abilities"`
	GenericTypeParams []GenericTypeParam  `json:"generic_type_params"`
	Fields            []ModuleStructField `json:"fields"`
}

type ABIInfo struct {
	Address          string            `json:"address"`
	Name             string            `json:"name"`
	Friends          []string          `json:"friends"`
	ExposedFunctions []ExposedFunction `json:"exposed_functions"`
	Structs          []ModuleStruct    `json:"structs"`
}

type AccountModuleInfo struct {
	Bytecode string  `json:"bytecode"`
	Abi      ABIInfo `json:"abi"`
}
