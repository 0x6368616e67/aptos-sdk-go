package types

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	buf, err := hex.DecodeString("5457B9493319D90188BF69187E9F8E8476258061341D86D6DB969A1E6C5FD7AD")
	assert.Equal(t, err, nil)
	privk := GenPrivKeyFromSeed(buf)
	pubk := privk.PubKey()
	addr := pubk.Address()
	t.Logf("addr:%s", addr.String())
	t.Logf("privKey:%s", privk.String())
	t.Logf("pubKey:%s", pubk.String())
	code := "b5e97db07fa0bd0e5598aa3643a9bc6f6693bddc1a9fec9e674a461eaa00b193eab474dcdfba9c26f57b33f8a93a2df1625586f231a9b7feca8bb3fc547d1e3f060000000000000002000000000000000000000000000000000000000000000000000000000000000104636f696e087472616e73666572010700000000000000000000000000000000000000000000000000000000000000010a6170746f735f636f696e094170746f73436f696e000220b5b3f30964642ff9406c092e89f320ace3ec8508e039b41f66adf7d466d52df9086400000000000000d00700000000000001000000000000001ade00630000000018"
	codeBuf, _ := hex.DecodeString(code)
	s, _ := privk.Sign(codeBuf)
	t.Logf("sign:%s", hex.EncodeToString(s))
}
