package onezero

import (
	"fmt"
	"testing"
)

const USER_NAME = "demo123"
const PASSWORD = "demo123"

func New() *Client {
	return NewClient(USER_NAME, PASSWORD, BASE_URL_PRODUCT)
}

// 测试获取token
func TestGetAccessToken(t *testing.T) {

	//请求
	resp, err := New().GetAccessToken()
	if err != nil {
		t.Errorf("request failed, msg[%v]", err)
	}
	fmt.Printf("accessToken: %+v\n", resp)
}
