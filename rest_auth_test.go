package onezero

import (
	"fmt"
	"testing"
)

func New() *Client {
	return NewClient(&OnezeroInitParams{USER_NAME, PASSWORD, BASE_URL_PRODUCT, REST_VERSION})
}

// 测试获取token
func TestGetAccessToken(t *testing.T) {

	//请求
	resp, err := New().GetAccessToken()
	if err != nil {
		t.Errorf("request failed, msg[%v]", err)
		return
	}
	fmt.Printf("accessToken: %+v\n", resp)
}
