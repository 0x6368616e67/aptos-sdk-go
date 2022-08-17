package v1

import (
	"fmt"

	"github.com/0x6368616e67/aptos-sdk-go/api/v1/general"
)

type MethodType uint16

const (
	MTHealthy MethodType = 0x1
	MTLedger  MethodType = 0x2
)

func Path(ant MethodType) string {
	p := ""
	switch ant {
	case MTHealthy:
		p = general.HealthyPath
	}
	return fmt.Sprintf("v1/%s", p)
}
