package types

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

type Transaction struct {
	Sender                  string                `json:"sender"`
	SequenceNumber          string                `json:"sequence_number"`
	MaxGasAmount            string                `json:"max_gas_amount"`
	GasUnitPrice            string                `json:"gas_unit_price"`
	ExpirationTimestampSecs string                `json:"expiration_timestamp_secs"`
	Payload                 TransactionPayload    `json:"payload"`
	Signature               *TransactionSignature `json:"signature,omitempty"`
	SecondarySigners        []string              `json:"secondary_signers"`
}
