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

func (cli *Client) requestJSON(ctx context.Context, msg interface{}) (rsp *rpcMessage, err error) {
	respBody, err := cli.conn.postJSON(ctx, msg)
	if err != nil {
		return
	}
	defer respBody.Close()

	rsp = &rpcMessage{}
	if err = json.NewDecoder(respBody).Decode(rsp); err != nil {
		return
	}
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
