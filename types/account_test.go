package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	account := NewAccountWithHexSeed("47F5F31C1E9D8C7F36A977904D2DE255C18BB9D9DD4F3EC6F28440473584C608")
	err := account.SyncSequence()
	assert.Equal(t, err, nil)
	t.Logf("%s:%+v", account.Address().String(), account)
}
