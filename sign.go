package aiguoclient_v2

import (
	"crypto/md5"
	"encoding/hex"
)

//内部函数
/*
	appID: 分配的渠道号
	appKey: 分配的秘钥
*/
func (cli *Client) genSign(body string) string {
	rawStr := body + cli.AppKey
	return calMd5(rawStr)
}

func calMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
