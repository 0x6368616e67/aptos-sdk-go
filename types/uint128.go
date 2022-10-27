package types

import (
	"errors"
	"math/big"

	"github.com/coming-chat/lcs"
)

func ReverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

type Uint128 struct{ *big.Int }

func (u Uint128) MarshalLCS(e *lcs.Encoder) error {
	if u.Sign() == -1 {
		return errors.New("Invalid U128")
	}
	bytes := u.Bytes()
	if len(bytes) > 16 {
		return errors.New("Invalid U128")
	}
	ReverseBytes(bytes)
	result := [16]byte{}
	copy(result[:], bytes)
	return e.EncodeFixedBytes(result[:])
}

func (u *Uint128) UnmarshalLCS(d *lcs.Decoder) error {
	bytes, err := d.DecodeFixedBytes(16)
	if err != nil {
		return err
	}
	ReverseBytes(bytes)
	u.Int = big.NewInt(0).SetBytes(bytes)
	return nil
}
