package onezero

import (
	"crypto/tls"
	"errors"
)

// 获取Token
// https://documentation-external.onezero.com:4434/RestAPI/onezero-rest-api.html#section/Authentication
func (cli *Client) GetAccessToken() (*AccessTokenResponse, error) {

	reqPath := "api/token"
	rawURL := cli.BaseURL + reqPath

	//返回值会放到这里
	result := AccessTokenResponse{}

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetFormData(map[string]string{
			"grant_type":   "password",
			"username":     cli.UserName,
			"password":     cli.Password,
			"rest_version": REST_VERSION,
		}).
		SetHeaders(getHeaders()).
		SetResult(&result). // or SetResult(AuthSuccess{}).
		Post(rawURL)

	if err != nil {
		return nil, err
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
