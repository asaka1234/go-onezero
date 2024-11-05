package onezero

import (
	"crypto/tls"
	"fmt"
)

// 获取指定margin account的持仓列表
// https://documentation-external.onezero.com:4434/RestAPI/onezero-rest-api.html#operation/GET-margin-account-id-positions
func (cli *Client) GetMarginAccountPositionList(auth string, marginAccountId int) (*MarginAccountPositionResponse, error) {

	reqPath := "api/rest/margin-account/%d/positions"
	rawURL := cli.BaseURL + fmt.Sprintf(reqPath, marginAccountId)

	//返回值会放到这里
	var result MarginAccountPositionResponse

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getAuthHeaders(auth)).
		SetResult(&result).
		Get(rawURL)

	if err != nil {
		return nil, err
	}

	fmt.Printf("==wsx====%s\n", string(resp.Body()))

	return &result, err
}
