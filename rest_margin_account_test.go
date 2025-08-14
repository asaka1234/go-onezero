package onezero

import (
	"fmt"
	"testing"
)

func TestGetMarginAccountPositionList(t *testing.T) {

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9zaWQiOiJiYWU3OTkyNi1mZGUwLTRjYzUtYjI5YS0wMDMwNjk4M2JiZWEiLCJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9uYW1lIjoiY3B0YnpfQWJvb2siLCJSZXN0VmVyc2lvbiI6IjEuMjAiLCJleHAiOjE3NDAwMjIyODh9.DHuOB2i_IHsCLUih51wKrlBcBqQhNc7XkAO2tt-r2rM"

	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9zaWQiOiI4YTBhMjNmMS1iOGMyLTQxODEtOTkzOC1mNmE5Nzk5YWNmOTMiLCJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9uYW1lIjoiY3B0YnoiLCJSZXN0VmVyc2lvbiI6IjEuMjAiLCJleHAiOjE3Mzk5Njk4MjR9.mZvj-whKUG1Bokk3sRR605U_wB5-MKrCsSoMl0dhntI"

	//请求
	resp, err := New().GetMarginAccountPositionList(token, 14)
	if err != nil {
		t.Errorf("request failed, msg[%v]", err)
		return
	}

	fmt.Printf("==wsx====%+v\n", resp)
}
