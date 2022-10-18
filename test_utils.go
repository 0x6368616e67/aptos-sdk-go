package aptos

import (
	"fmt"
	"net/http"
)

var (
	faucetURLFmt = "https://tap.devnet.prod.gcp.aptosdev.com/mint?address=%s&amount=%d"
)

func faucet(addr string, amount uint64) (err error) {
	faucetURL := fmt.Sprintf(faucetURLFmt, addr, amount)
	fmt.Printf("Get faucet:%s \n", faucetURL)
	_, err = http.Post(faucetURL, "", nil)
	return err

}
