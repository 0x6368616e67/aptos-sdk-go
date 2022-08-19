package aptos

const (
	Devnet string = "https://fullnode.devnet.aptoslabs.com"
)

var (
	Endpoint string
)

func init() {
	Endpoint = Devnet
}
