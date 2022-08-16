package aptos

import (
	"context"
	"encoding/json"
	"net/url"
)

// Client represent a RPC Client
type Client struct {
	conn *httpConn
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
	return
}

func (cli *Client) requestJSONRPC(ctx context.Context, msg interface{}) (rsp *rpcMessage, err error) {
	respBody, err := cli.conn.doRequest(ctx, msg)
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
