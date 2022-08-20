package aptos

const (
	Devnet string = "https://fullnode.devnet.aptoslabs.com"
)

var (
	Endpoint     string
	MaxGasAmount = 2000
	GasUnitPrice = 1
)

func init() {
	Endpoint = Devnet
}
