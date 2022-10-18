package framework

import (
	"github.com/0x6368616e67/aptos-sdk-go/types"
)

type AptosAccount struct {
	sequenceNumber    uint64
	authenticationKey string
	privateKey        types.PrivKey
}

func NewAptosAccount(privKey *types.PrivKey, authAddr types.Address) *AptosAccount {
	if privKey == nil {
		key := types.GenPrivKey()
		privKey = &key
	}
	acc := &AptosAccount{
		privateKey:        *privKey,
		sequenceNumber:    0,
		authenticationKey: authAddr.String(),
	}
	return acc
}

func (acc *AptosAccount) Address() types.Address {
	return acc.privateKey.PubKey().Address()
}
