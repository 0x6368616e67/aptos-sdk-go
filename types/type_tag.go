package types

import (
	"errors"
	"strings"

	"github.com/coming-chat/lcs"
)

func init() {
	lcs.RegisterEnum(
		(*TypeTag)(nil),

		TypeTagBool{},
		TypeTagU8{},
		TypeTagU64{},
		TypeTagU128{},
		TypeTagAddress{},
		TypeTagSigner{},
		TypeTagVector{},
		TypeTagStruct{},
	)
}

type TypeTag interface{}

type TypeTagBool struct{}
type TypeTagU8 struct{}
type TypeTagU64 struct{}
type TypeTagU128 struct{}
type TypeTagAddress struct{}
type TypeTagSigner struct{}
type TypeTagVector struct {
	Value TypeTag `lcs:"value"`
}
type TypeTagStruct struct {
	Address    Address   `lcs:"address"`
	ModuleName string    `lcs:"module_name"`
	Name       string    `lcs:"name"`
	TypeArgs   []TypeTag `lcs:"type_args"`
}

func NewTypeTagStructFromString(tag string) (*TypeTagStruct, error) {
	parts := strings.Split(tag, "::")
	if len(parts) != 3 {
		return nil, errors.New("Invalid struct tag string")
	}
	addr := HexToAddress(parts[0])
	return &TypeTagStruct{
		Address:    addr,
		ModuleName: parts[1],
		Name:       parts[2],
	}, nil
}
