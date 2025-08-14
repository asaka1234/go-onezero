package onezero

import (
	"crypto/tls"
	"errors"
	"fmt"
)

// 获取Token
// https://documentation-external.onezero.com:4434/RestAPI/onezero-rest-api.html#section/Authentication
func (cli *Client) GetAccessToken() (*AccessTokenResponse, error) {

	reqPath := "api/token"
	rawURL := cli.Params.BaseURL + reqPath

	//返回值会放到这里
	result := AccessTokenResponse{}

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetDebug(cli.debugMode).
		SetFormData(map[string]string{
			"grant_type":   "password",
			"username":     cli.Params.UserName,
			"password":     cli.Params.Password,
			"rest_version": cli.Params.RestVersion,
		}).
		SetHeaders(getHeaders()).
		SetResult(&result). // or SetResult(AuthSuccess{}).
		Post(rawURL)

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

	if result.AccessToken == "" {
		//说明没有拿到,此时要去看错误
		/*
			errResult := CommonErrorResponse{}
			rawBody := string(resp.Body())
			json.Unmarshal([]byte(rawBody), &errResult)
		*/
		return nil, errors.New(string(resp.Body()))
	}

	return &result, nil
}
