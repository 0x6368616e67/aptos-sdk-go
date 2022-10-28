package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/coming-chat/lcs"
	"golang.org/x/crypto/sha3"
)

const (
	RAW_TRANSACTION_SALT           = "APTOS::RawTransaction"
	RAW_TRANSACTION_WITH_DATA_SALT = "APTOS::RawTransactionWithData"
)

func init() {
	lcs.RegisterEnum(
		(*RawTransactionPayload)(nil),

		RawTransactionPayloadScript{},
		RawTransactionPayloadModuleBundle{},
		RawTransactionPayloadEntryFunction{},
	)

	lcs.RegisterEnum(
		(*RawTransactionArgument)(nil),

		RawTransactionArgumentU8{},
		RawTransactionArgumentU64{},
		RawTransactionArgumentU128{},
		RawTransactionArgumentAddress{},
		RawTransactionArgumentU8Vector{},
		RawTransactionArgumentBool{},
	)

	lcs.RegisterEnum(
		(*RawRawTransactionWithData)(nil),
	)
}

type RawTransactionPayloadScript struct {
	Code   []byte                   `lcs:"code"`
	TyArgs []TypeTag                `lcs:"ty_args"`
	Args   []RawTransactionArgument `lcs:"args"`
}

type RawTransactionPayloadEntryFunction struct {
	ModuleName   ModuleID  `lcs:"module_name"`
	FunctionName string    `lcs:"function_name"`
	TyArgs       []TypeTag `lcs:"ty_args"`
	Args         [][]byte  `lcs:"args"`
}

type RawTransactionPayloadModuleBundle struct{}

type Module struct {
	Code []byte `lcs:"code"`
}

type ModuleID struct {
	Address Address `lcs:"address"`
	Name    string  `lcs:"name"`
}

func ParseModuleAndFunctionName(function string) (*ModuleID, string, error) {
	parts := strings.Split(function, "::")
	if len(parts) != 3 {
		return nil, "", errors.New("invalid function literal")
	}
	addr := HexToAddress(parts[0])
	return &ModuleID{
		addr,
		parts[1],
	}, parts[2], nil
}

type RawTransactionArgument interface{}

type RawTransactionArgumentU8 struct {
	Value uint8 `lcs:"value"`
}
type RawTransactionArgumentU64 struct {
	Value uint64 `lcs:"value"`
}
type RawTransactionArgumentU128 struct {
	Value Uint128 `lcs:"value"`
}
type RawTransactionArgumentAddress struct {
	Value Address `lcs:"value"`
}
type RawTransactionArgumentU8Vector struct {
	Value []uint8 `lcs:"value"`
}
type RawTransactionArgumentBool struct {
	Value bool `lcs:"value"`
}

type RawRawTransactionWithData interface{}

type RawTransactionPayload interface{}
type RawTransaction struct {
	Sender                  Address               `lcs:"sender"`
	SequenceNumber          uint64                `lcs:"sequence_number"`
	Payload                 RawTransactionPayload `lcs:"payload"`
	MaxGasAmount            uint64                `lcs:"max_gas_amount"`
	GasUnitPrice            uint64                `lcs:"gas_unit_price"`
	ExpirationTimestampSecs uint64                `lcs:"expiration_timestamp_secs"`
	ChainID                 uint8                 `lcs:"chain_id"`
}

type TransactionPayload struct {
	Type          string        `json:"type"`
	Function      string        `json:"function"`
	TypeArguments []string      `json:"type_arguments"`
	Arguments     []interface{} `json:"arguments"`
}

type TransactionSignature struct {
	Type      string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type InnerTransaction struct {
	Sender                  string             `json:"sender"`
	SequenceNumber          string             `json:"sequence_number"`
	MaxGasAmount            string             `json:"max_gas_amount"`
	GasUnitPrice            string             `json:"gas_unit_price"`
	ExpirationTimestampSecs string             `json:"expiration_timestamp_secs"`
	Payload                 TransactionPayload `json:"payload"`
}

type Transaction struct {
	InnerTransaction
	Signature        *TransactionSignature `json:"signature"`
	SecondarySigners *[]string             `json:"secondary_signers,omitempty"`
}

func (tx *Transaction) ToRawTransaction(payload RawTransactionPayload) *RawTransaction {
	rawTx := &RawTransaction{}
	rawTx.Sender = HexToAddress(tx.Sender)
	rawTx.SequenceNumber, _ = strconv.ParseUint(tx.SequenceNumber, 10, 64)
	rawTx.Payload = payload
	rawTx.MaxGasAmount, _ = strconv.ParseUint(tx.MaxGasAmount, 10, 64)
	rawTx.GasUnitPrice, _ = strconv.ParseUint(tx.GasUnitPrice, 10, 64)
	rawTx.ExpirationTimestampSecs, _ = strconv.ParseUint(tx.ExpirationTimestampSecs, 10, 64)
	chainID, _ := strconv.ParseUint(tx.MaxGasAmount, 10, 8)
	rawTx.ChainID = uint8(chainID)
	return rawTx
}

func (tx *Transaction) EncodeToBCS(payload RawTransactionPayload) (data []byte, err error) {
	rawTx := tx.ToRawTransaction(payload)
	rawTx.ChainID = 35
	fmt.Printf("rawTx:%+v \n", rawTx)
	prefixBytes := sha3.Sum256([]byte(RAW_TRANSACTION_SALT))
	data, err = lcs.Marshal(rawTx)
	data = append(prefixBytes[:], data...)
	return
}
