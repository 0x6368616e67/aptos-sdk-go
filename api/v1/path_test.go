package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	rp, m := Path(MTTransactionByHash)
	assert.Equal(t, m, "GET")
	assert.Equal(t, rp, "v1/transactions/by_hash/{txn_hash}")

	rp, m = Path(MTLedger)
	assert.Equal(t, m, "GET")
	assert.Equal(t, rp, "v1/")
}
