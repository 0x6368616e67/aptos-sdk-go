package aptos

import (
	"context"
	"testing"

	"github.com/0x6368616e67/aptos-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

const (
	devnet = "https://fullnode.devnet.aptoslabs.com"
)

func TestHealthy(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	err = cli.Healthy(context.Background(), 1)
	assert.Equal(t, err, nil, "Healthy error")
}

func TestLedgerInfo(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	info, err := cli.LedgerInfo(context.Background())
	t.Logf("info:%+v \n", info)
	assert.Equal(t, err, nil, "LedgerInfo error")
	assert.Equal(t, info.ChainID, 23)
}

func TestGetAccount(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	info, err := cli.GetAccount(context.Background(), "0x5c96ae24729caa96958df32f0c8ca715494d738e943b14961541e477b133ea9c", 0)
	t.Logf("info:%+v \n", info)
	assert.Equal(t, err, nil, "GetAccount error")
	assert.Greater(t, len(info.AuthenticationKey), 0)
}

func TestGetAccountResource(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	info, err := cli.GetAccountResource(context.Background(), "0x81873855df80aae3e1468e4d47d85f4be2126df25574d29b40cb57be01a93c1c", 0)
	t.Logf("info:%+v \n", info)
	assert.Equal(t, err, nil, "GetAccount error")
	assert.Greater(t, len(info), 0)
}

func TestGetAccountResourceWithType(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	info, err := cli.GetAccountResourceWithType(context.Background(), "0x81873855df80aae3e1468e4d47d85f4be2126df25574d29b40cb57be01a93c1c", "0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>", 0)
	t.Logf("info:%+v \n", info)
	assert.Equal(t, err, nil, "GetAccountResourceWithType error")
	assert.Greater(t, len(info.Data), 0)
}

func TestGetAccountModule(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	info, err := cli.GetAccountModule(context.Background(), "0x1", 0)
	t.Logf("info:%+v \n", info)
	assert.Equal(t, err, nil, "GetAccountModule error")
	assert.Greater(t, len(info), 0)
}

func TestGetAccountModuleWithName(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	info, err := cli.GetAccountModuleWithName(context.Background(), "0x1", "acl", 0)
	t.Logf("info:%+v \n", info)
	assert.Equal(t, err, nil, "GetAccountModule error")
	assert.Greater(t, len(info.Bytecode), 0)
}

func TestGetBlock(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	info, err := cli.GetBlock(context.Background(), 2042680, true)
	t.Logf("info:%+v \n", info)
	assert.Equal(t, err, nil, "GetBlock error")
	assert.Equal(t, info.BlockHeight, "2042680")

	_, err = cli.GetBlock(context.Background(), 20426801232323, true)
	assert.NotEqual(t, err, nil)
	t.Logf("err:%+v \n", err)
}

func TestGetEvent(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	infos, err := cli.GetEvent(context.Background(), "0x010000000000000030b26281346f628a63c8c0adb7042f991e6702141dcada1968171460c99bbe60", 0, 0)
	t.Logf("info:%+v \n", infos)
	assert.Equal(t, err, nil, "GetEvent error")
	assert.Greater(t, len(infos), 0)
}

func TestGetEventWithHandler(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	infos, err := cli.GetEventWithHandler(context.Background(), "0x30b26281346f628a63c8c0adb7042f991e6702141dcada1968171460c99bbe60", "0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>", "deposit_events", 0)
	t.Logf("info:%+v \n", infos)
	assert.Equal(t, err, nil, "GetEventWithHandler error")
	assert.Greater(t, len(infos), 0)
}

func TestGetTransactions(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	infos, err := cli.GetTransactions(context.Background(), 0, 2)
	t.Logf("info:%+v \n", infos)
	assert.Equal(t, err, nil, "GetTransactions error")
	assert.Greater(t, len(infos), 0)
}

func TestGetTransactionsOfAccount(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	infos, err := cli.GetTransactionsOfAccount(context.Background(), "0x3f25e8f59cacfda5a18152e3ddd08926969416f466c38db95ed00efcd318b971", 0, 2)
	t.Logf("info:%+v \n", infos)
	assert.Equal(t, err, nil, "GetTransactionsOfAccount error")
	assert.Greater(t, len(infos), 0)
}

func TestGetTransactionByHash(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	info, err := cli.GetTransactionByHash(context.Background(), "0x11a6129c84ec02e0b8e192a61e58bf234583863299397d2d9aaece72de167dfb")
	t.Logf("info:%+v \n", info)
	assert.Equal(t, err, nil, "GetTransactionByHash error")
	assert.Equal(t, info.Hash, "0x11a6129c84ec02e0b8e192a61e58bf234583863299397d2d9aaece72de167dfb")
}

func TestGetTransactionByVersion(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	info, err := cli.GetTransactionByVersion(context.Background(), 25125141)
	t.Logf("info:%+v \n", info)
	assert.Equal(t, err, nil, "GetTransactionByVersion error")
	assert.Equal(t, info.Hash, "0x11a6129c84ec02e0b8e192a61e58bf234583863299397d2d9aaece72de167dfb")
}

func TestGetTransactionEncoding(t *testing.T) {
	cli, err := DialContext(context.Background(), devnet)
	assert.Equal(t, err, nil, "DialContext error")
	var args []interface{}
	args = append(args, "0x3aa1e96803500900ed3bdd8cc779fefe7e88aafd015a65b2aa5eb39bda2e1f47")
	args = append(args, "20000")
	tx := types.Transaction{
		InnerTransaction: types.InnerTransaction{
			Sender:                  "0x3dc12eb3816bdf291b28e544cf88c1fb647d613ff63cb464bfb59fb2bf941ec6",
			SequenceNumber:          "7635",
			MaxGasAmount:            "2000",
			GasUnitPrice:            "1",
			ExpirationTimestampSecs: "1660903177",
			Payload: types.TransactionPayload{
				Type:          "entry_function_payload",
				Function:      "0x1::aptos_coin::mint",
				Arguments:     args,
				TypeArguments: make([]string, 0),
			},
		},
		SecondarySigners: nil,
	}
	code, err := cli.GetTransactionEncoding(context.Background(), &tx)
	t.Logf("code:%+v \n", code)
	assert.Equal(t, err, nil, "GetTransactionEncoding error")
}
