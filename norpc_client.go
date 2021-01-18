package plugin

import "errors"

type NORPCClient struct {
}

func newNORPCClient(c *Client) (*NORPCClient, error)             { return &NORPCClient{}, nil }
func (c *NORPCClient) Dispense(name string) (interface{}, error) { return nil, nil }
func (c *NORPCClient) Ping() error                               { return errors.New("NOPING") }
func (c *NORPCClient) Close() error                              { return nil }
