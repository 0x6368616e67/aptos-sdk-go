package aptos

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/0x6368616e67/aptos-sdk-go/framework"
	"github.com/0x6368616e67/aptos-sdk-go/types"
)

type Account struct {
	privateKey types.PrivKey
	sequence   uint64
	cli        *Client
}

func NewAccount() *Account {
	acc := &Account{
		sequence: 0,
	}

	acc.privateKey = types.GenPrivKey()
	var err error
	if acc.cli, err = Dial(Endpoint); err != nil {
		return nil
	}
	return acc
}

func NewAccountWithPrivateKey(key string) *Account {
	key = strings.TrimPrefix(key, "0x")
	acc := &Account{}
	var err error
	if acc.privateKey, err = types.GenPrivKeyFromHex(key); err != nil {
		return nil
	}

	if acc.cli, err = Dial(Endpoint); err != nil {
		return nil
	}
	return acc
}

func NewAccountWithHexSeed(seed string) *Account {
	acc := &Account{
		sequence: 0,
	}
	seedBuf, err := hex.DecodeString(seed)
	if err != nil {
		return nil
	}
	acc.privateKey = types.GenPrivKeyFromSeed(seedBuf)

	if acc.cli, err = Dial(Endpoint); err != nil {
		return nil
	}
	return acc
}

func (acc *Account) PrivateKey() string {
	return acc.privateKey.String()
}

func (acc *Account) Reconnect() (err error) {
	if acc.cli, err = Dial(Endpoint); err != nil {
		return err
	}
	return nil
}

func (acc *Account) Sign(msg []byte) ([]byte, error) {
	return acc.privateKey.Sign(msg)
}

func (acc *Account) Sequence() uint64 {
	return acc.sequence
}

func (acc *Account) SendTransaction(tx *types.Transaction) (hash string, err error) {
	err = acc.SignTx(tx)
	if err != nil {
		return
	}
	rst, err := acc.cli.SubmitTransaction(context.Background(), tx)
	if err != nil {
		return
	}
	hash = rst.Hash
	return
}

func (acc *Account) SimulateTransaction(tx *types.Transaction) (err error) {
	err = acc.SignTx(tx)
	if err != nil {
		return
	}
	tx.Signature.Signature = "0x" + strings.Repeat("0", len(tx.Signature.Signature)-2)
	rst, err := acc.cli.SimulateTransaction(context.Background(), tx)
	if err != nil {
		return
	}
	if len(rst) == 0 {
		return errors.New("result is empty")
	}
	if !rst[0].Success {
		return errors.New(rst[0].VMStatus)
	}
	return nil
}

func (acc *Account) SignTx(tx *types.Transaction) (err error) {
	code, err := acc.cli.GetTransactionEncoding(context.Background(), tx)
	if err != nil {
		return err
	}
	codeBuf, err := hex.DecodeString(code[2:])
	if err != nil {
		return err
	}
	sign, err := acc.Sign(codeBuf)
	if err != nil {
		return err
	}
	signHex := hex.EncodeToString(sign)
	signHex = "0x" + signHex
	tx.Signature = &types.TransactionSignature{
		Type:      "ed25519_signature",
		PublicKey: acc.privateKey.PubKey().String(),
		Signature: signHex,
	}
	return nil
}

func (acc *Account) Address() types.Address {
	return acc.privateKey.PubKey().Address()
}

func (acc *Account) SyncSequence() error {
	info, err := acc.cli.GetAccount(context.Background(), acc.Address().String(), 0)
	if err != nil {
		return err
	}
	acc.sequence, err = strconv.ParseUint(info.SequenceNumber, 10, 64)
	if err != nil {
		return err
	}
	return nil
}

func (acc *Account) Transfer(to types.Address, amount uint64) (hash string, err error) {
	err = acc.SyncSequence()
	if err != nil {
		return
	}

	var args []interface{}
	args = append(args, to.String())
	args = append(args, strconv.FormatUint(amount, 10))
	tx := types.Transaction{
		InnerTransaction: types.InnerTransaction{
			Sender:                  acc.Address().String(),
			SequenceNumber:          strconv.FormatUint(acc.sequence, 10),
			MaxGasAmount:            strconv.FormatUint(uint64(MaxGasAmount), 10),
			GasUnitPrice:            strconv.FormatUint(uint64(GasUnitPrice), 10),
			ExpirationTimestampSecs: strconv.FormatUint(uint64(time.Now().Unix()+10*60), 10),
			Payload: types.TransactionPayload{
				Type:          "entry_function_payload",
				Function:      "0x1::coin::transfer",
				Arguments:     args,
				TypeArguments: []string{"0x1::aptos_coin::AptosCoin"},
			},
		},
		SecondarySigners: nil,
	}
	hash, err = acc.SendTransaction(&tx)
	return
}

func (acc *Account) Balance() (balance uint64, err error) {

	info, err := acc.cli.GetAccountResourceWithType(context.Background(), acc.Address().String(), "0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>", 0)
	if err != nil {
		return
	}
	if info == nil || info.Data == nil {
		return
	}
	type Coin_0x1_aptos_coin_AptosCoin struct {
		Coin struct {
			Value string `json:"value"`
		} `json:"coin"`
		DepositEvents struct {
			Counter string `json:"counter"`
			GUID    struct {
				ID struct {
					Addr        string `json:"addr"`
					CreationNum string `json:"creation_num"`
				} `json:"id"`
			} `json:"guid"`
		} `json:"deposit_events"`
		Frozen         bool `json:"frozen"`
		WithdrawEvents struct {
			Counter string `json:"counter"`
			GUID    struct {
				ID struct {
					Addr        string `json:"addr"`
					CreationNum string `json:"creation_num"`
				} `json:"id"`
			} `json:"guid"`
		} `json:"withdraw_events"`
	}
	coin := Coin_0x1_aptos_coin_AptosCoin{}
	err = json.Unmarshal(info.Data, &coin)
	if err != nil {
		return
	}
	balance, err = strconv.ParseUint(coin.Coin.Value, 10, 64)
	return
}

func (acc *Account) CreateAptosAccount() (account *framework.AptosAccount, err error) {
	account = framework.NewAptosAccount(nil, acc.Address())
	err = acc.SyncSequence()
	if err != nil {
		return
	}

	var args []interface{}
	args = append(args, account.Address().String())
	tx := types.Transaction{
		InnerTransaction: types.InnerTransaction{
			Sender:                  acc.Address().String(),
			SequenceNumber:          strconv.FormatUint(acc.sequence, 10),
			MaxGasAmount:            strconv.FormatUint(uint64(MaxGasAmount), 10),
			GasUnitPrice:            strconv.FormatUint(uint64(GasUnitPrice), 10),
			ExpirationTimestampSecs: strconv.FormatUint(uint64(time.Now().Unix()+10*60), 10),
			Payload: types.TransactionPayload{
				Type:          "entry_function_payload",
				Function:      "0x1::aptos_account::create_account",
				Arguments:     args,
				TypeArguments: make([]string, 0),
			},
		},
		SecondarySigners: nil,
	}
	_, err = acc.SendTransaction(&tx)
	return
}
