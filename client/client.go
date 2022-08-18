package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"reflect"

	v1 "github.com/0x6368616e67/aptos-sdk-go/api/v1"
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

func (cli *Client) request(ctx context.Context, method v1.MethodType, param interface{}, result interface{}) error {
	if result != nil && reflect.TypeOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("result parameter must be pointer or nil interface: %v", result)
	}
	urlpath := v1.Path(method)
	resp, err := cli.requestGet(ctx, urlpath, param)
	if err != nil {
		if e, ok := err.(HTTPError); ok {
			errmsg := ErrorMsg{}
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
	err = cli.request(ctx, v1.MTAccountResourceWithType, param, &info)
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
	err = cli.request(ctx, v1.MTAccountModuleWithName, param, &info)
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
	err = cli.request(ctx, v1.MTBlock, param, &info)
	if err != nil {
		return nil, err
	}
	return
}
