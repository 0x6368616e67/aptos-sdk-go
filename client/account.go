package client

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/0x6368616e67/aptos-sdk-go"
	"github.com/0x6368616e67/aptos-sdk-go/types"
)

type Account struct {
	privateKey types.PrivKey
	sequence   uint64
	cli        *Client
}

func NewAccount() *Account {
	acc := &Account{
		sequence: 0,
	}

	acc.privateKey = types.GenPrivKey()
	var err error
	if acc.cli, err = Dial(aptos.Endpoint); err != nil {
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
	acc.privateKey = types.GenPrivKeyFromSeed(seedBuf)

	if acc.cli, err = Dial(aptos.Endpoint); err != nil {
		return nil
	}
	return acc
}

func (acc *Account) Sign(msg []byte) ([]byte, error) {
	return acc.privateKey.Sign(msg)
}

func (acc *Account) SendTransaction(tx *types.Transaction) (hash string, err error) {
	err = acc.SignTx(tx)
	if err != nil {
		return
	}
	fmt.Printf("tx:%+v\n", tx)
	rst, err := acc.cli.SubmitTransaction(context.Background(), tx)
	if err != nil {
		return
	}
	hash = rst.Hash
	return
}

func (acc *Account) SignTx(tx *types.Transaction) (err error) {
	code, err := acc.cli.GetTransactionEncoding(context.Background(), tx)
	if err != nil {
		return err
	}
	fmt.Printf("code:%+v", code)
	codeBuf, err := hex.DecodeString(code[2:])
	if err != nil {
		return err
	}
	sign, err := acc.Sign(codeBuf)
	if err != nil {
		return err
	}
	signHex := hex.EncodeToString(sign)
	signHex = "0x" + signHex
	tx.Signature = &types.TransactionSignature{
		Type:      "ed25519_signature",
		PublicKey: acc.privateKey.PubKey().String(),
		Signature: signHex,
	}
	return nil
}

func (acc *Account) Address() types.Address {
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
