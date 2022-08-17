package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var endpoints = []string{
	"https://aptoslabs.com/",
	"http://aptoslabs.com/",
}

type Request struct {
	Hello string `json:"hello"`
}

func TestDoRequest(t *testing.T) {
	for _, ep := range endpoints {
		req := &Request{"World"}
		c := newHTTPConn(ep)
		_, err := c.postJSON(context.Background(), req)
		assert.NotEqual(t, err, nil, "should be a 404")

	}
}
