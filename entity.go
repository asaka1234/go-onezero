package aiguoclient

type Request struct {
	AppId    string `json:"appId"`
	SignType string `json:"signType"`
	Sign     string `json:"sign"`
	Data     string `json:"data"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"`
}

// ---------------------------------------------------
type CreateOrderOptionsV2 struct {
	MerchantOrderNo  string `json:"merchantOrderNo" comment:"商家订单号"  valid:"Required"`
	ParcelName       string `json:"parcelName" comment:"快递包裹物名称"  valid:"Required"`
	ParcelNum        int    `json:"parcelNum" comment:"包裹数量"  valid:"Required"`
	SenderProvince   string `json:"senderProvince" comment:"发件人省份"  valid:"Required"`
	SenderCity       string `json:"senderCity" comment:"发件人城市"  valid:"Required"`
	SenderDistrict   string `json:"senderDistrict" comment:"发件人地区"  valid:"Required"`
	SenderAddress    string `json:"senderAddress" comment:"发件人地址"  valid:"Required"`
	SenderLinkPerson string `json:"senderLinkPerson" comment:"发件人名称"  valid:"Required"`
	SenderLinkTel    string `json:"senderLinkTel" comment:"发件人手机号"  valid:"Required"`

	ReceiveProvince   string `json:"receiveProvince" comment:"收件人省份"  valid:"Required"`
	ReceiveCity       string `json:"receiveCity" comment:"收件人城市"  valid:"Required"`
	ReceiveDistrict   string `json:"receiveDistrict" comment:"收件人地区"  valid:"Required"`
	ReceiveAddress    string `json:"receiveAddress" comment:"收件人地址"  valid:"Required"`
	ReceiveLinkPerson string `json:"receiveLinkPerson" comment:"收件人名称"  valid:"Required"`
	ReceiveLinkTel    string `json:"receiveLinkTel" comment:"收件人手机号"  valid:"Required"`

	EstimateWeight   float64 `json:"estimateWeight" comment:"预估重量，重量范围1-50"  valid:"Required"`
	AppointmentStart string  `json:"appointmentStart" comment:"预约开始时间 2024-06-01 09:00:00"  valid:"Required"`
	AppointmentEnd   string  `json:"appointmentEnd" comment:"预约结束时间 2024-06-01 10:00:00"  valid:"Required"`
	UserRemark       string  `json:"userRemark,omitempty" comment:"用户备注"`
}

//----取消-----

type CancelOrderOptionsV2 struct {
	MerchantOrderNo string `json:"merchantOrderNo" comment:"商家订单号"`
	CancelReason    string `json:"cancelReason" comment:"取消原因"`
}

// -------- 查询订单 ----------------------------
type QueryOrderOptionsV2 struct {
	MerchantOrderNo string `json:"merchantOrderNo" comment:"商家订单号"`
}

type QueryOrderResponseV2 struct {
	MerchantOrderNo string `json:"merchantOrderNo" comment:"商家订单号"`
	OrderNo         string `json:"orderNo"`

	ParcelName string `json:"parcelName" comment:"快递包裹物名称"`
	WaybillNo  string `json:"waybillNo,omitempty" comment:"运单号"`

	CancelReason     string `json:"cancelReason,omitempty"`
	ExpressCompanyNo string `json:"expressCompanyNo,omitempty"`
	ExpressStatus    string `json:"expressStatus"` //CREATED

	ReceiveCountry    string `json:"receiveCountry"`
	ReceiveAddress    string `json:"receiveAddress"`
	ReceiveCity       string `json:"receiveCity"`
	ReceiveDistrict   string `json:"receiveDistrict"`
	ReceiveLinkPerson string `json:"receiveLinkPerson"`
	ReceiveLinkTel    string `json:"receiveLinkTel"`
	SenderCountry     string `json:"senderCountry"`
	ReceiveProvince   string `json:"receiveProvince"`
	SenderAddress     string `json:"senderAddress"`
	SenderCity        string `json:"senderCity"`
	SenderDistrict    string `json:"senderDistrict"`
	SenderLinkPerson  string `json:"senderLinkPerson"`
	SenderLinkTel     string `json:"senderLinkTel"`
	SenderProvince    string `json:"senderProvince"`

	EstimateWeight   float64 `json:"estimateWeight" comment:"预估重量，重量范围1-50"`
	AppointmentStart string  `json:"appointmentStart" comment:"预约开始时间 2024-06-01 09:00:00"`
	AppointmentEnd   string  `json:"appointmentEnd" comment:"预约结束时间 2024-06-01 10:00:00"`
}

// -------------------------------------------
// 轨迹追踪
type TraceOrderOptionsV2 struct {
	WaybillNo string `json:"waybillNo" comment:"运单号"`
}

// 每一条轨迹item
type TraceOrderResponseItemV2 struct {
	Time        string `json:"time"`
	Description string `json:"description"`
}

// ------ 回调返回 ---------------
type CallbackResponseV2 struct {
	ExpressCompanyNo string `json:"expressCompanyNo" comment:"快递公司编码"`
	MerchantOrderNo  string `json:"merchantOrderNo" comment:"商家订单号"`
	NotifyTime       string `json:"notifyTime" comment:"通知时间"`
	OrderNo          string `json:"orderNo" comment:"平台订单号"`
	Status           string `json:"status" comment:"当前订单状态"`
	WaybillNo        string `json:"waybillNo" comment:"	运单号"`
	ExpressWaybills  []struct {
		TrackingNo string `json:"trackingNo" comment:"	面单号"`
		WaybillNo  string `json:"waybillNo" comment:"运单号"`
	} `json:"expressWaybills"`
	CourierName   string  `json:"courierName,omitempty" comment:"揽收快递员名称,订单状态COLLECTING时有值"`
	CourierTel    string  `json:"courierTel,omitempty" comment:"揽收员电话,订单状态COLLECTING时有值"`
	ActualWeight  float64 `json:"actualWeight,omitempty" comment:"实际重量，订单状态TRANSIT时有值"`
	BillingWeight float64 `json:"billingWeight,omitempty" comment:"计费重量，订单状态TRANSIT时有值"`
	CancelReason  string  `json:"cancelReason,omitempty" comment:"取消原因"`
}
