package general

const (
	HealthyPath = "-/healthy"
)

type HealthyReq struct {
	Duration uint32 `param:"duration_secs"`
}

type HealthyRsp struct {
	Message string `json:"message"`
}
