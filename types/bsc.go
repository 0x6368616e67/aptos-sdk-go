package types

import (
	"github.com/coming-chat/lcs"
)

func EncodeBCS[T bool | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | string | Uint128](t T) []byte {
	s, _ := lcs.Marshal(t)
	return s
}
