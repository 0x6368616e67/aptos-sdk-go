package client

import (
	"testing"

	v1 "github.com/0x6368616e67/aptos-sdk-go/api/v1"
	"github.com/stretchr/testify/assert"
)

func TestEncodeURLParam(t *testing.T) {
	req := v1.HealthyReq{
		Duration: 123,
	}

	assert.Equal(t, "duration_secs=123", encodeURLParam(req))
}

func TestEndodePathParam(t *testing.T) {
	req := v1.AccountReq{
		Address: "0xabc",
	}

	assert.Equal(t, "accounts/0xabc", endodePathParam("accounts/{address}", req))
}

func TestEndodeURLPath(t *testing.T) {
	req := v1.AccountReq{
		Address:       "0xabc",
		LedgerVersion: 123,
	}

	assert.Equal(t, "accounts/0xabc?ledger_version=123", endodeURLPath("accounts/{address}", req))

	req = v1.AccountReq{
		Address: "0xabc",
	}

	assert.Equal(t, "accounts/0xabc", endodeURLPath("accounts/{address}", req))
}
