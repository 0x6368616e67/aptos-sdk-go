package types

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	buf, err := hex.DecodeString("5457B9493319D90188BF69187E9F8E8476258061341D86D6DB969A1E6C5FD7AD")
	assert.Equal(t, err, nil)
	privk := GenPrivKeyFrom32Bytes(buf)
	pubk := privk.PubKey()
	addr := pubk.Address()
	t.Logf("addr:%s", addr.String())
	t.Logf("privKey:%s", privk.String())
	t.Logf("pubKey:%s", pubk.String())
}
