# aptos-sdk-go
Aptos SDK for Golang

[![Go Build status](https://github.com/0x6368616e67/aptos-sdk-go/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/0x6368616e67/aptos-sdk-go/actions/workflows/build.yml)[![Test status](https://github.com/0x6368616e67/aptos-sdk-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/0x6368616e67/aptos-sdk-go/actions/workflows/ci.yml) [![SDK Documentation](https://img.shields.io/badge/SDK-Documentation-blue)](https://pkg.go.dev/github.com/0x6368616e67/aptos-sdk-go) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/0x6368616e67/aptos-sdk-go/blob/main/LICENSE)

`aptos-sdk-go` is a golang sdk for [Aptos](https://aptoslabs.com/). Which contains 
all [RPC API](https://fullnode.mainnet.aptoslabs.com/v1/spec#/) with `Client` and some operation 
for `Account` object such as `Transfer`, `Balance` etc...

## Getting started

Add SDK Dependencies

    $ go get -u  github.com/0x6368616e67/aptos-sdk-go

A "Alice and Bob" example  demonstrate two user Alice and Bob.
each of them faucet 10000 coin first. and then Alice transfer 
5000 to Bob.

Here is the code

    package main

    import (
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
        alice := aptos.NewAccount()
        bob := aptos.NewAccount()
        fmt.Printf("faucet first \n")
        err := faucet(alice.Address().String(), 10000)
        if err != nil {
            panic(err.Error())
        }
        err = faucet(bob.Address().String(), 10000)
        if err != nil {
            panic(err.Error())
        }

        fmt.Printf("wait 10 second ...")
        time.Sleep(10 * time.Second) // wait for stat
        aliceBalance, err := alice.Balance()
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Alice balance:%d\n", aliceBalance)

        bobBalance, err := bob.Balance()
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Bob balance:%d\n", bobBalance)

        hash, err := alice.Transfer(bob.Address(), 5000)
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Alice transfer %d to bob with hash:%s\n", 5000, hash)
        fmt.Printf("====================================================\n")
        fmt.Printf("wait 10 second ...")
        time.Sleep(10 * time.Second) // wait for stat
        aliceBalance, err = alice.Balance()
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Alice balance:%d\n", aliceBalance)

        bobBalance, err = bob.Balance()
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Bob balance:%d\n", bobBalance)

    }

when run we got:


    alice_and_bob % go run main.go 
    faucet first 
    Get faucet:https://tap.devnet.prod.gcp.aptosdev.com/mint?address=0x62c0ce0d9d2d6bea852abd6f4feb6f88dea0a4a4eabc9edf295a4949f3f47870&amount=10000 
    Get faucet:https://tap.devnet.prod.gcp.aptosdev.com/mint?address=0x0b74ceb77162d79c5d61a546e8a2e47c29e35f0fced9e961e8c7a919802ed0de&amount=10000 
    wait 10 second ...
    Alice balance:10000
    Bob balance:10000
    Alice transfer 5000 to bob with hash:0x6566c80b00cb66a79cbe153ed96827fc62a8b57d4675dd260813edf644fe3939
    ====================================================
    wait 10 second ...
    Alice balance:4999
    Bob balance:15000

Then in the explorer , we can see what happend [0x6566c80b00cb66a79cbe153ed96827fc62a8b57d4675dd260813edf644fe3939](https://explorer.devnet.aptos.dev/txn/2604671) (as devnet will be reload, this version for the hash will change, use your own please.)



## Support RPC Status
### Accounts: Access to account resources and modules

- [x] Get account
- [x] Get account resources
- [x] Get account modules
- [x] Get specific account resource 
- [x] Get specific account module
### Blocks: Access to blocks

- [x] Get blocks by height

### Events: Access to events

- [x] Get events by event key
- [x] Get events by event handle

### General: General information

- [ ] ~~Show OpenAPI explorer~~
- [x] Check basic node health
- [x] Get ledger info

### Tables: Access to tables

- [ ] Get table item

### Transactions: Access to transactions

- [x] Get transactions
- [x] Submit transaction
- [x] Get transaction by hash
- [x] Get transaction by version
- [x] Get account transactions
- [x] Simulate transaction
- [x] Encode submission


## TODO

- [ ] Add wait status for sending a transaction
- [ ] Add non native token operation for 0x3::token
- [ ] Add support for x25519
- [ ] Add more examples
- [ ] Add more document
- [ ] Add support for tables
- [ ] Add support multi signer

