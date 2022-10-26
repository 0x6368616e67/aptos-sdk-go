package aptos

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"reflect"

	v1 "github.com/0x6368616e67/aptos-sdk-go/api/v1"
	"github.com/0x6368616e67/aptos-sdk-go/types"
)

// Client represent a RPC Client
type Client struct {
	conn   *httpConn
	rpcurl string
}

// Dial connects a client to the given URL.
func Dial(rpcurl string) (*Client, error) {
	return DialContext(context.Background(), rpcurl)
}

// Dial connects a client to the given URL and given context
func DialContext(ctx context.Context, rpcurl string) (client *Client, err error) {
	_, err = url.Parse(rpcurl)
	if err != nil {
		return nil, err
	}

	client = &Client{}
	client.conn = newHTTPConn(rpcurl)
	client.rpcurl = rpcurl
	return
}

func (cli *Client) requestGet(ctx context.Context, urlpath string, msg interface{}) (rsp []byte, err error) {
	respBody, err := cli.conn.get(ctx, urlpath, msg)
	if err != nil {
		return
	}
	defer respBody.Close()

	rsp, err = ioutil.ReadAll(respBody)
	return
}

func (cli *Client) requestPost(ctx context.Context, urlpath string, msg interface{}) (rsp []byte, err error) {
	respBody, err := cli.conn.postJSON(ctx, urlpath, msg)
	if err != nil {
		return
	}
	defer respBody.Close()

	rsp, err = ioutil.ReadAll(respBody)
	return
}

func (cli *Client) request(ctx context.Context, method v1.MethodType, param interface{}, result interface{}) (err error) {
	if result != nil && reflect.TypeOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("result parameter must be pointer or nil interface: %v", result)
	}
	var resp []byte
	urlpath, httpMethod := v1.Path(method)
	if httpMethod == "GET" {
		resp, err = cli.requestGet(ctx, urlpath, param)
	} else if httpMethod == "POST" {
		resp, err = cli.requestPost(ctx, urlpath, param)
	}
	if err != nil {
		if e, ok := err.(HTTPError); ok {
			errmsg := v1.ErrorMsg{}
			json.Unmarshal(e.Body, &errmsg)
			return errmsg
		}
		return err
	}
	json.Unmarshal(resp, result)
	return nil
}

func (cli *Client) Healthy(ctx context.Context, duration uint32) error {
	param := v1.HealthyReq{
		Duration: duration,
	}
	rsp := v1.HealthyRsp{}
	err := cli.request(ctx, v1.MTHealthy, param, &rsp)
	return err
}

func (cli *Client) LedgerInfo(ctx context.Context) (info *v1.LedgerInfo, err error) {
	rsp := v1.LedgerRsp{}
	err = cli.request(ctx, v1.MTLedger, nil, &rsp)
	if err != nil {
		return nil, err
	}
	info = &rsp.LedgerInfo
	return
}

func (cli *Client) GetAccount(ctx context.Context, address string, ledger uint64) (info *v1.AccountInfo, err error) {
	param := v1.AccountReq{
		Address:       address,
		LedgerVersion: ledger,
	}
	info = &v1.AccountInfo{}
	err = cli.request(ctx, v1.MTAccount, param, info)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetAccountResource(ctx context.Context, address string, ledger uint64) (infos []*v1.AccountResourceInfo, err error) {
	param := v1.AccountResourceReq{
		Address:       address,
		LedgerVersion: ledger,
	}
	infos = make([]*v1.AccountResourceInfo, 1)
	err = cli.request(ctx, v1.MTAccountResource, param, &infos)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetAccountResourceWithType(ctx context.Context, address string, resType string, ledger uint64) (info *v1.AccountResourceInfo, err error) {
	param := v1.AccountResourceWithTypeReq{
		Address:       address,
		LedgerVersion: ledger,
		Type:          resType,
	}
	info = &v1.AccountResourceInfo{}
	err = cli.request(ctx, v1.MTAccountResourceWithType, param, info)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetAccountModule(ctx context.Context, address string, ledger uint64) (infos []*v1.AccountModuleInfo, err error) {
	param := v1.AccountModuleReq{
		Address:       address,
		LedgerVersion: ledger,
	}
	infos = make([]*v1.AccountModuleInfo, 1)
	err = cli.request(ctx, v1.MTAccountModule, param, &infos)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetAccountModuleWithName(ctx context.Context, address string, name string, ledger uint64) (info *v1.AccountModuleInfo, err error) {
	param := v1.AccountModuleWithNameReq{
		Address:       address,
		LedgerVersion: ledger,
		Name:          name,
	}
	info = &v1.AccountModuleInfo{}
	err = cli.request(ctx, v1.MTAccountModuleWithName, param, info)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetBlock(ctx context.Context, height uint64, withTransactions bool) (info *v1.BlockInfo, err error) {
	param := v1.BlockReq{
		BlockHeight:      height,
		WithTransactions: withTransactions,
	}
	info = &v1.BlockInfo{}
	err = cli.request(ctx, v1.MTBlock, param, info)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetEvent(ctx context.Context, key string, limit uint16, start uint64) (infos []*v1.EventInfo, err error) {
	param := v1.EventReq{
		EventKey: key,
		Limit:    limit,
		Start:    start,
	}
	infos = make([]*v1.EventInfo, 1)
	err = cli.request(ctx, v1.MTEvent, param, &infos)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetEventWithHandler(ctx context.Context, address string, handler string, field string, limit uint16) (infos []*v1.EventInfo, err error) {
	param := v1.EventWithHandlerReq{
		Address: address,
		Handler: handler,
		Filed:   field,
		Limit:   0,
	}
	infos = make([]*v1.EventInfo, 1)
	err = cli.request(ctx, v1.MTEventWithHandler, param, &infos)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetTransactions(ctx context.Context, start uint64, limit uint16) (infos []*v1.TransactionInfo, err error) {
	param := v1.TransactionReq{
		Start: start,
		Limit: limit,
	}
	infos = make([]*v1.TransactionInfo, 1)
	err = cli.request(ctx, v1.MTTransaction, param, &infos)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetTransactionByHash(ctx context.Context, hash string) (info *v1.TransactionInfo, err error) {
	param := v1.TransactionByHashReq{
		Hash: hash,
	}
	info = &v1.TransactionInfo{}
	err = cli.request(ctx, v1.MTTransactionByHash, param, info)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetTransactionByVersion(ctx context.Context, version uint64) (info *v1.TransactionInfo, err error) {
	param := v1.TransactionByVersionReq{
		Version: version,
	}
	info = &v1.TransactionInfo{}
	err = cli.request(ctx, v1.MTTransactionByVersion, param, info)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetTransactionsOfAccount(ctx context.Context, address string, start uint64, limit uint16) (infos []*v1.TransactionInfo, err error) {
	param := v1.TransactionOfAccountReq{
		Address: address,
		Start:   start,
		Limit:   limit,
	}
	infos = make([]*v1.TransactionInfo, 1)
	err = cli.request(ctx, v1.MTTransactionOfAccount, param, &infos)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) GetTransactionEncoding(ctx context.Context, tx *types.Transaction) (code string, err error) {
	err = cli.request(ctx, v1.MTTransactionEncoding, tx, &code)
	return
}

func (cli *Client) SubmitTransaction(ctx context.Context, tx *types.Transaction) (result *v1.TransactionInfo, err error) {
	result = &v1.TransactionInfo{}
	err = cli.request(ctx, v1.MTTransactionSubmit, tx, result)
	return
}

func (cli *Client) SimulateTransaction(ctx context.Context, tx *types.Transaction) (result []*v1.TransactionInfo, err error) {
	result = make([]*v1.TransactionInfo, 1)
	err = cli.request(ctx, v1.MTTransactionSimulate, tx, &result)
	return
}
