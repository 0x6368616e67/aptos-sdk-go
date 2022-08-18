package v1

import (
	"fmt"
)

type MethodType uint16

const (
	MTHealthy MethodType = 0x1
	MTLedger  MethodType = 0x2
	MTAccount MethodType = 0x3
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
	}
	return fmt.Sprintf("v1/%s", p)
}
