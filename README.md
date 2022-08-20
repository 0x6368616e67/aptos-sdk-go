# aptos-sdk-go
Aptos SDK for Golang

[![Go Build status](https://github.com/0x6368616e67/aptos-sdk-go/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/0x6368616e67/aptos-sdk-go/actions/workflows/build.yml)[![Test status](https://github.com/0x6368616e67/aptos-sdk-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/0x6368616e67/aptos-sdk-go/actions/workflows/ci.yml) [![SDK Documentation](https://img.shields.io/badge/SDK-Documentation-blue)](https://pkg.go.dev/github.com/0x6368616e67/aptos-sdk-go) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/0x6368616e67/aptos-sdk-go/blob/main/LICENSE)


## RPC Status
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
- [ ] Add support for tables
- [ ] Add support multi signer

