package v1

import (
	"fmt"
	"strings"
)

type MethodType uint16

const (
	MTHealthy                 MethodType = 0x1
	MTLedger                  MethodType = 0x2
	MTAccount                 MethodType = 0x3
	MTAccountResource         MethodType = 0x4
	MTAccountModule           MethodType = 0x5
	MTAccountResourceWithType MethodType = 0x6
	MTAccountModuleWithName   MethodType = 0x7
	MTTransaction             MethodType = 0xb
	MTTransactionByHash       MethodType = 0xc
	MTTransactionByVersion    MethodType = 0xd
	MTTransactionOfAccount    MethodType = 0xe

	MTTransactionEncoding   MethodType = 0x11
	MTTransactionSimulate   MethodType = 0x12
	MTTransactionSubmit     MethodType = 0x13
	MTBlockByHeight         MethodType = 0x14
	MTBlockByVersion        MethodType = 0x15
	MTEventByCreationNumber MethodType = 0x16
	MTEventByEventHandler   MethodType = 0x17
	MTEstimateGasPrice      MethodType = 0x18
)

func Path(ant MethodType) (rawpath string, method string) {
	p := ""
	switch ant {
	case MTHealthy:
		p = HealthyPath
	case MTLedger:
		p = LedgerPath
	case MTAccount:
		p = AccountPath
	case MTAccountResource:
		p = AccountResourcePath
	case MTAccountModule:
		p = AccountModulePath
	case MTAccountResourceWithType:
		p = AccountResourceWithTypePath
	case MTAccountModuleWithName:
		p = AccountModuleWithNamePath
	case MTBlockByHeight:
		p = BlockByHeightPath
	case MTBlockByVersion:
		p = BlockByVersionPath
	case MTEventByCreationNumber:
		p = EventByCreationNumberPath
	case MTEventByEventHandler:
		p = EventByEventHandlerPath
	case MTTransaction:
		p = TransactionPath
	case MTTransactionByHash:
		p = TransactionByHashPath
	case MTTransactionByVersion:
		p = TransactionByVersionPath
	case MTTransactionOfAccount:
		p = TransactionOfAccountPath
	case MTTransactionEncoding:
		p = TransactionEncodingPath
	case MTTransactionSimulate:
		p = TransactionSimulatePath
	case MTTransactionSubmit:
		p = TransactionSubmitPath
	case MTEstimateGasPrice:
		p = EstimateGasPricePath
	}
	items := strings.Split(p, "@")
	if len(items) == 1 {
		rawpath = p
		method = "GET"
	} else {
		method = items[0]
		rawpath = strings.Join(items[1:], "@")
	}
	rawpath = fmt.Sprintf("v1/%s", rawpath)
	return
}
