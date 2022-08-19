package client

import (
	"testing"

	"github.com/0x6368616e67/aptos-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	account := NewAccountWithHexSeed("47F5F31C1E9D8C7F36A977904D2DE255C18BB9D9DD4F3EC6F28440473584C608")
	err := account.SyncSequence()
	assert.Equal(t, err, nil)
	t.Logf("%s:%+v", account.Address().String(), account)
}

func TestAccountSubmitTransaction(t *testing.T) {
	account := NewAccountWithHexSeed("47F5F31C1E9D8C7F36A977904D2DE255C18BB9D9DD4F3EC6F28440473584C608")
	var args []interface{}
	args = append(args, "0x3aa1e96803500900ed3bdd8cc779fefe7e88aafd015a65b2aa5eb39bda2e1f47")
	args = append(args, "20000")
	tx := types.Transaction{
		Sender:                  account.Address().String(),
		SequenceNumber:          "1",
		MaxGasAmount:            "2000",
		GasUnitPrice:            "1",
		ExpirationTimestampSecs: "1660903177",
		Payload: types.TransactionPayload{
			Type:          "entry_function_payload",
			Function:      "0x1::aptos_coin::mint",
			Arguments:     args,
			TypeArguments: make([]string, 0),
		},
		SecondarySigners: make([]string, 0),
	}
	hash, err := account.SendTransaction(&tx)
	assert.Equal(t, err, nil)
	t.Logf("%s", hash)
}
