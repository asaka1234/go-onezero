package aiguoclient_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/mitchellh/mapstructure"
	"github.com/xxjwxc/public/errors"
	"io/ioutil"
	"log"
	"net/http"
)

// 下单
// 返回的第一个参数是：成功时的平台订单号
func (cli *Client) AddOrder(bodyParam CreateOrderOptionsV2) (string, error) {

	reqPath := "api/createOrder"
	rawURL := cli.BaseURL + reqPath

	// param check
	valid := validation.Validation{}
	b, err := valid.Valid(&bodyParam)
	if err != nil {
		return "", err
	}
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		return "", errors.New("参数错误")
	}

	//data数据
	paramJSON, _ := json.Marshal(bodyParam)

	//构造最终的请求body
	requestFinal := cli.genRequest(string(paramJSON))

	req, _ := http.NewRequest("POST", rawURL, bytes.NewReader(requestFinal))
	req.Header.Add("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respByte, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		//请求失败了
		return "", errors.New("code!=200")
	}

	var urlResp Response
	json.Unmarshal(respByte, &urlResp)
	if urlResp.Status != 200 {
		return "", errors.New(fmt.Sprintf("status=%d, msg=%s", urlResp.Status, urlResp.Msg))
	}

	return urlResp.Data.(string), nil

}

//-------------------------------------------
//运单轨迹
/*
	waybillNo     运单号
*/
func (cli *Client) ListTrace(bodyParam TraceOrderOptionsV2) ([]TraceItem, error) {

	reqPath := "/api/order/trace"
	rawURL := cli.BaseURL + reqPath

	//data数据
	paramJSON, _ := json.Marshal(bodyParam)

	//构造最终的请求body
	requestFinal := cli.genRequest(string(paramJSON))

	//-----------
	req, _ := http.NewRequest("POST", rawURL, bytes.NewReader(requestFinal))
	req.Header.Add("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respByte, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		//请求失败了
		return nil, errors.New("code!=200")
	}

	var urlResp Response
	json.Unmarshal(respByte, &urlResp)
	if urlResp.Status != 200 {
		return nil, errors.New(fmt.Sprintf("status=%d, msg=%s", urlResp.Status, urlResp.Msg))
	}

	//--------解析data------------------------
	var result []TraceOrderResponseItemV2
	err = mapstructure.Decode(urlResp.Data.([]interface{}), &result)
	if err != nil {
		return nil, err //errors.New(fmt.Sprintf("status=%d, msg=%s", urlResp.Status, urlResp.Msg))
	}

	//----adapter v1----------------------------
	if len(result) > 0 {
		f := make([]TraceItem, len(result))
		for _, item := range result {
			f = append(f, TraceItem{
				Time: item.Time,
				Desc: item.Description,
			})
		}
		return f, nil
	}
	//-----------------------------------------------
	return nil, nil
}

//-------------------------------------------

// 取消
func (cli *Client) CancelOrder(bodyParam CancelOrderOptionsV2) error {

	reqPath := "api/cancelOrder"
	rawURL := cli.BaseURL + reqPath

	// param check
	valid := validation.Validation{}
	b, err := valid.Valid(&bodyParam)
	if err != nil {
		return err
	}
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		return errors.New("参数错误")
	}

	//body数据
	paramJSON, _ := json.Marshal(bodyParam)

	//构造最终的请求body
	requestFinal := cli.genRequest(string(paramJSON))

	//-----------
	req, _ := http.NewRequest("POST", rawURL, bytes.NewReader(requestFinal))
	req.Header.Add("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respByte, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		//请求失败了
		return errors.New("code!=200")
	}

	var urlResp Response
	json.Unmarshal(respByte, &urlResp)
	if urlResp.Status != 200 {
		return errors.New(fmt.Sprintf("status=%d, msg=%s", urlResp.Status, urlResp.Msg))
	}

	return nil
}

//--------------------------------

// 查询订单
func (cli *Client) QueryOrder(bodyParam QueryOrderOptionsV2) (QueryOrderResponseV2, error) {

	reqPath := "api/queryOrder"
	rawURL := cli.BaseURL + reqPath

	//body数据
	paramJSON, _ := json.Marshal(bodyParam)

	//构造最终的请求body
	requestFinal := cli.genRequest(string(paramJSON))

	//-----------
	req, _ := http.NewRequest("POST", rawURL, bytes.NewReader(requestFinal))
	req.Header.Add("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return QueryOrderResponseV2{}, err
	}
	defer resp.Body.Close()

	respByte, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		//请求失败了
		return QueryOrderResponseV2{}, errors.New("code!=200")
	}

	var urlResp Response
	json.Unmarshal(respByte, &urlResp)
	if urlResp.Status != 200 {
		return QueryOrderResponseV2{}, errors.New(fmt.Sprintf("status=%d, msg=%s", urlResp.Status, urlResp.Msg))
	}

	//--------解析data------------------------
	var result QueryOrderResponseV2
	err = mapstructure.Decode(urlResp.Data.(map[string]interface{}), &result)
	if err != nil {
		return QueryOrderResponseV2{}, err //errors.New(fmt.Sprintf("status=%d, msg=%s", urlResp.Status, urlResp.Msg))
	}

	return result, nil
}

//-------------------------------

// AiGuoCallback 统一回调入口
func AiGuoCallback(w http.ResponseWriter, r *http.Request) (CallbackResponseV2, error) {
	var succeed = 1
	defer func() {
		if succeed == 0 {
			w.WriteHeader(500)
		}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		succeed = 0
		return CallbackResponseV2{}, err
	}

	//把数据转到struct里
	var resp CallbackResponseV2
	err = json.Unmarshal(body, &resp)
	if err != nil {
		succeed = 0
		return CallbackResponseV2{}, err
	}

	return resp, nil
}

//----------------------tool----------------------------------

func (cli *Client) genRequest(data string) []byte {
	//计算签名
	sign := cli.genSign(data)
	request := Request{
		AppId:    fmt.Sprintf("%d", cli.AppID),
		SignType: "MD5",
		Sign:     sign,
		Data:     data,
	}
	requestFinal, _ := json.Marshal(request)
	return requestFinal
}
