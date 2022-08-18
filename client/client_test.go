package client

import (
	"context"
	"testing"

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
