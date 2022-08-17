package general

import "github.com/0x6368616e67/aptos-sdk-go/types"

const (
	HealthyPath = "-/healthy"
)

type HealthyReq struct {
	Duration uint32 `param:"duration_secs"`
}

type HealthyRsp struct {
	types.JSONRsp
}
