package types

import (
	"context"
	"encoding/hex"
	"strconv"

	"github.com/0x6368616e67/aptos-sdk-go"
	"github.com/0x6368616e67/aptos-sdk-go/client"
)

type Account struct {
	privateKey PrivKey
	sequence   uint64
	cli        *client.Client
}

func NewAccount() *Account {
	acc := &Account{
		sequence: 0,
	}

	acc.privateKey = GenPrivKey()
	var err error
	if acc.cli, err = client.Dial(aptos.Endpoint); err != nil {
		return nil
	}
	return acc
}

func NewAccountWithHexSeed(seed string) *Account {
	acc := &Account{
		sequence: 0,
	}
	seedBuf, err := hex.DecodeString(seed)
	if err != nil {
		return nil
	}
	acc.privateKey = GenPrivKeyFromSeed(seedBuf)

	if acc.cli, err = client.Dial(aptos.Endpoint); err != nil {
		return nil
	}
	return acc
}

func (acc *Account) Sign(msg []byte) ([]byte, error) {
	return acc.privateKey.Sign(msg)
}

func (acc *Account) Address() Address {
	return acc.privateKey.PubKey().Address()
}

func (acc *Account) SyncSequence() error {
	info, err := acc.cli.GetAccount(context.Background(), acc.Address().String(), 0)
	if err != nil {
		return err
	}
	acc.sequence, err = strconv.ParseUint(info.SequenceNumber, 10, 64)
	if err != nil {
		return err
	}
	return nil
}
