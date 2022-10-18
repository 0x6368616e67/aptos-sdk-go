package aptos

const (
	Devnet string = "https://fullnode.devnet.aptoslabs.com"
)

var (
	Endpoint     string
	MaxGasAmount = 1000
	GasUnitPrice = 100
)

func init() {
	Endpoint = Devnet
}
