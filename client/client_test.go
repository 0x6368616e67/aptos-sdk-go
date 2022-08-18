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
