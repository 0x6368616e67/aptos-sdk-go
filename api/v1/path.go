package v1

import (
	"fmt"
)

type MethodType uint16

const (
	MTHealthy                 MethodType = 0x1
	MTLedger                  MethodType = 0x2
	MTAccount                 MethodType = 0x3
	MTAccountResource         MethodType = 0x4
	MTAccountModule           MethodType = 0x5
	MTAccountResourceWithType MethodType = 0x6
)

func Path(ant MethodType) string {
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
	}
	return fmt.Sprintf("v1/%s", p)
}
