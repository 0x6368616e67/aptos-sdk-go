package aptos

const (
	Devnet  string = "https://fullnode.devnet.aptoslabs.com"
	Mainnet string = "https://fullnode.mainnet.aptoslabs.com"
)

var (
	Endpoint     string
	MaxGasAmount = 10000
	GasUnitPrice = 100
)

func init() {
	Endpoint = Devnet
}
