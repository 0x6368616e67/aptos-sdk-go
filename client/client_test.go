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
