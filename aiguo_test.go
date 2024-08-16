package aiguoclient

import (
	"fmt"
	_ "image/gif"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const APP_ID = 20240715
const APP_KEY = "ee67d727001047879d367a9d9c0d9acf"

func New() *Client {
	return NewClient(APP_ID, APP_KEY, BASE_URL_DEV)
}

// 2024081616171302098
// 测试下单
func TestAddOrder(t *testing.T) {

	//参数
	param := CreateOrderOptionsV2{

		MerchantOrderNo: "m1234edc",
		ParcelName:      "大衣",
		ParcelNum:       1,

		SenderProvince:   "山东",
		SenderCity:       "青岛",
		SenderDistrict:   "崂山区",
		SenderAddress:    "中山路23号",
		SenderLinkPerson: "收货人姓名",
		SenderLinkTel:    "19821232111",

		ReceiveProvince:   "山东",
		ReceiveCity:       "青岛",
		ReceiveDistrict:   "崂山区",
		ReceiveAddress:    "中山路24号",
		ReceiveLinkPerson: "收货人姓名",
		ReceiveLinkTel:    "19821232009",

		EstimateWeight:   1.5,
		AppointmentStart: "2024-08-17 09:00:00",
		AppointmentEnd:   "2024-08-17 10:00:00",
		UserRemark:       "",
	}

	//请求
	platOrderNo, err := New().AddOrder(param)
	if err != nil {
		t.Errorf("request failed, msg[%v]", err)
	}
	fmt.Printf("platOrderNo: %s\n", platOrderNo)
}

// 测试查询订单详情
func TestQueryOrder(t *testing.T) {

	param := QueryOrderOptionsV2{
		MerchantOrderNo: "m1234edc", //商户订单号
	}

	//运单轨迹
	detail, err := New().QueryOrder(param)
	if err != nil {
		t.Errorf("request failed, msg[%v]", err)
	}
	fmt.Printf("detail:%+v\n", detail)
}

// 测试获取轨迹
func TestListTrace(t *testing.T) {

	//参数
	reqData := TraceOrderOptionsV2{
		WaybillNo: "SF3123968234911", //运单号
	}

	//运单轨迹
	traceList, err := New().ListTrace(reqData)
	if err != nil {
		t.Errorf("request failed, msg[%v]", err)
	}
	fmt.Println(traceList)
}

// 测试取消
func TestCancelOrder(t *testing.T) {

	//参数
	param := CancelOrderOptionsV2{
		MerchantOrderNo: "12345", //商户订单号
		CancelReason:    "地址写错了",
	}

	//请求
	err := New().CancelOrder(param)
	if err != nil {
		t.Errorf("request failed, msg[%v]", err)
	}
	fmt.Printf("取消成功\n")
}

// 分配快递员成功，开始去取件 : 此时可以查出【快递员的name和手机号】 --> 取间中
// 取件成功 : 此时可以拿到【真实重量 和 数量】 --> 运输中
func TestCallback(t *testing.T) {

	//取件中
	//content := "{\"courierName\":\"42409443\",\"courierTel\":\"18610525235\",\"expressCompanyNo\":\"SF\",\"expressWaybills\":[{\"trackingNo\":\"SF3123968234911\",\"waybillNo\":\"SF3123968234911\"}],\"merchantOrderNo\":\"TH202408154229\",\"notifyTime\":\"2024-08-15 17:48:22\",\"orderNo\":\"202408151743317094\",\"status\":\"COLLECTING\",\"waybillNo\":\"SF3123968234911\"} "
	//已揽收
	//content := "{\"expressCompanyNo\":\"SF\",\"expressWaybills\":[{\"trackingNo\":\"SF3123968234911\",\"waybillNo\":\"SF3123968234911\"}],\"merchantOrderNo\":\"TH202408154229\",\"notifyTime\":\"2024-08-15 20:20:07\",\"orderNo\":\"202408151743317094\",\"status\":\"COLLECTED\",\"waybillNo\":\"SF3123968234911\"}  "
	//运输中
	//content := "{\"actualWeight\":9.40,\"billingWeight\":9.50,\"expressCompanyNo\":\"SF\",\"expressFee\":17.00,\"expressWaybills\":[{\"trackingNo\":\"SF3123968234911\",\"waybillNo\":\"SF3123968234911\"}],\"merchantOrderNo\":\"TH202408154229\",\"notifyTime\":\"2024-08-15 21:32:31\",\"orderNo\":\"202408151743317094\",\"status\":\"TRANSIT\",\"waybillNo\":\"SF3123968234911\"}  "
	//已签收
	content := "{\"actualWeight\":9.40,\"billingWeight\":9.50,\"expressCompanyNo\":\"SF\",\"expressWaybills\":[{\"trackingNo\":\"SF3123968234911\",\"waybillNo\":\"SF3123968234911\"}],\"merchantOrderNo\":\"TH202408154229\",\"notifyTime\":\"2024-08-16 15:26:11\",\"orderNo\":\"202408151743317094\",\"status\":\"SIGNED\",\"waybillNo\":\"SF3123968234911\"}  "

	w := httptest.NewRecorder()
	payload := strings.NewReader(content)
	req := httptest.NewRequest(http.MethodPut, "http://127.0.0.1:8128/notify", payload)

	//调用
	resp, err := AiGuoCallback(w, req)
	if err != nil {
		t.Errorf("request failed:%s", err.Error())
		return
	}
	fmt.Printf("%+v\n", resp)
}
