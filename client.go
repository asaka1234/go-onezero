package onezero

import "github.com/go-resty/resty/v2"

type Client struct {
	Params *OnezeroInitParams

	debugMode bool
	ryClient  *resty.Client
}

func NewClient(params *OnezeroInitParams) *Client {
	return &Client{
		Params:    params,
		debugMode: false,
		ryClient:  resty.New(), //client实例
	}
}

func (cli *Client) SetDebugModel(debugMode bool) {
	cli.debugMode = debugMode
}
