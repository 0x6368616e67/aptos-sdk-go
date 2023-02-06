package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	aptos "github.com/0x6368616e67/aptos-sdk-go"
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

func main() {
	cli, err := aptos.DialContext(context.Background(), aptos.Devnet)
	if err != nil {
		panic(err.Error())
	}
	alice := aptos.NewAccount()
	bob := aptos.NewAccount()
	fmt.Printf("faucet first \n")
	err = faucet(alice.Address().String(), 1000000)
	if err != nil {
		panic(err.Error())
	}
	err = faucet(bob.Address().String(), 1000000)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("wait 10 second ...")
	time.Sleep(10 * time.Second) // wait for stat
	aliceBalance, err := alice.Balance(cli)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Alice balance:%d\n", aliceBalance)

	bobBalance, err := bob.Balance(cli)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Bob balance:%d\n", bobBalance)

	hash, err := alice.Transfer(cli, bob.Address(), 5000)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Alice transfer %d to bob with hash:%s\n", 5000, hash)
	fmt.Printf("====================================================\n")
	fmt.Printf("wait 10 second ...\n")
	time.Sleep(10 * time.Second) // wait for stat
	aliceBalance, err = alice.Balance(cli)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Alice balance:%d\n", aliceBalance)

	bobBalance, err = bob.Balance(cli)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Bob balance:%d\n", bobBalance)

}
