package onezero

import "github.com/go-resty/resty/v2"

type Client struct {
	UserName string
	Password string
	BaseURL  string
	ryClient *resty.Client
}

func NewClient(userName string, password string, baseURL string) *Client {
	return &Client{
		UserName: userName,
		Password: password,
		BaseURL:  baseURL,
		ryClient: resty.New(), //client实例
	}
}
