package onezero

import (
	"crypto/tls"
	"fmt"
)

// 获取指定margin account的持仓列表
// https://documentation-external.onezero.com:4434/RestAPI/onezero-rest-api.html#operation/GET-margin-account-id-positions
func (cli *Client) GetMarginAccountPositionList(auth string, marginAccountId int) (*MarginAccountPositionResponse, error) {

	reqPath := "api/rest/margin-account/%d/positions"
	rawURL := cli.Params.BaseURL + fmt.Sprintf(reqPath, marginAccountId)

	//返回值会放到这里
	var result MarginAccountPositionResponse

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetDebug(cli.debugMode).
		SetHeaders(getAuthHeaders(auth)).
		SetResult(&result).
		Get(rawURL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	return &result, err
}
