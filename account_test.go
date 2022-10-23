package aptos

import (
	"strconv"
	"testing"
	"time"

	"github.com/0x6368616e67/aptos-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	account := NewAccountWithHexSeed("47F5F31C1E9D8C7F36A977904D2DE255C18BB9D9DD4F3EC6F28440473584C608")

	faucet(account.Address().String(), 1000000)
	time.Sleep(1 * time.Second)
	err := account.SyncSequence()
	assert.Equal(t, err, nil)
	t.Logf("%s:%+v", account.Address().String(), account)
}

func TestAccountSubmitTransaction(t *testing.T) {
	//account := NewAccountWithHexSeed("5457B9493319D90188BF69187E9F8E8476258061341D86D6DB969A1E6C5FD7AD")
	account := NewAccount()
	faucet(account.Address().String(), 1000000)
	time.Sleep(1 * time.Second)
	err := account.SyncSequence()
	assert.Equal(t, err, nil)
	var args []interface{}
	args = append(args, "0xb5b3f30964642ff9406c092e89f320ace3ec8508e039b41f66adf7d466d52df9")
	args = append(args, "100")
	tx := types.Transaction{
		InnerTransaction: types.InnerTransaction{
			Sender:                  account.Address().String(),
			SequenceNumber:          strconv.FormatUint(account.sequence, 10),
			MaxGasAmount:            "2000",
			GasUnitPrice:            "100",
			ExpirationTimestampSecs: strconv.FormatUint(uint64(time.Now().Unix()+600), 10),
			Payload: types.TransactionPayload{
				Type:          "entry_function_payload",
				Function:      "0x1::coin::transfer",
				Arguments:     args,
				TypeArguments: []string{"0x1::aptos_coin::AptosCoin"},
			},
		},
		SecondarySigners: nil,
	}
	hash, err := account.SendTransaction(&tx)
	assert.Equal(t, err, nil)
	t.Logf("%s", hash)
}

func TestAccountSubmitTransaction2(t *testing.T) {
	//account := NewAccountWithHexSeed("5457B9493319D90188BF69187E9F8E8476258061341D86D6DB969A1E6C5FD7AD")
	account := NewAccount()
	faucet(account.Address().String(), 1000000)
	time.Sleep(1 * time.Second)
	err := account.SyncSequence()
	assert.Equal(t, err, nil)
	var args []interface{}
	args = append(args, account.Address().String())
	args = append(args, "100")
	tx := types.Transaction{
		InnerTransaction: types.InnerTransaction{
			Sender:                  account.Address().String(),
			SequenceNumber:          strconv.FormatUint(account.sequence, 10),
			MaxGasAmount:            "2000",
			GasUnitPrice:            "1",
			ExpirationTimestampSecs: strconv.FormatUint(uint64(time.Now().Unix()+600), 10),
			Payload: types.TransactionPayload{
				Type:          "entry_function_payload",
				Function:      "0x1::aptos_coin::mint",
				Arguments:     args,
				TypeArguments: make([]string, 0),
			},
		},
		SecondarySigners: nil,
	}
	hash, err := account.SendTransaction(&tx)
	assert.Equal(t, err, nil)
	t.Logf("%s", hash)
}

func TestAccountTransfer(t *testing.T) {
	account := NewAccountWithHexSeed("5457B9493319D90188BF69187E9F8E8476258061341D86D6DB969A1E6C5FD7AD")
	faucet(account.Address().String(), 1000000)
	time.Sleep(1 * time.Second)
	hash, err := account.Transfer(types.HexToAddress("0xb5b3f30964642ff9406c092e89f320ace3ec8508e039b41f66adf7d466d52df9"), 1000)
	assert.Equal(t, err, nil)
	t.Logf("%s", hash)
}

func TestAccountBalance(t *testing.T) {
	account := NewAccountWithHexSeed("5457B9493319D90188BF69187E9F8E8476258061341D86D6DB969A1E6C5FD7AD")
	faucet(account.Address().String(), 1000000)
	time.Sleep(1 * time.Second)
	balance, err := account.Balance()
	assert.Equal(t, err, nil)
	t.Logf("balance:%d", balance)
}
