package aiguoclient_v2

const (
	CATEGORY_CLOTHES = 1 //	衣物
	CATEGORY_MOBILE  = 2 // 移动电话
	CATEGORY_BOOK    = 3 //	书籍
	CATEGORY_OTHER   = 4 //	其它
)

// 所有的默认类型配置
var CONF_CATEGORY_AIGO_CODE = map[int]string{
	CATEGORY_CLOTHES: "CLOTHES",
	CATEGORY_MOBILE:  "MOBILE",
	CATEGORY_BOOK:    "BOOK",
	CATEGORY_OTHER:   "OTHER",
}

const (
	NOTIFY_CATEGORY_SENT              = 1 //	aiguo 给 顺丰 下单成功
	NOTIFY_CATEGORY_RECEIPTING        = 2 //    分配快递员成功，开始去取件
	NOTIFY_CATEGORY_GOT               = 3 //	取件成功
	NOTIFY_CATEGORY_SIGNEDFROMEXPRESS = 4 //	收件人签收成功
	NOTIFY_CATEGORY_CANCELLED         = 5 //	取消物流
	NOTIFY_CATEGORY_AIGUO_CANCEL      = 6 //    aighuo取消
)

const (
	SEND_FAILURE     = 1 //物流下单失败
	SEND_SUCCESS     = 2 //物流下单成功
	COLLECTING       = 3 //揽收中
	COLLECTED        = 4 //已揽收
	TRANSIT          = 5 //运输中
	SIGNED           = 6 //已签收
	CANCELED         = 7 //已取消
	CANCELED_INVALID = 8 //订单作废
	CANCELED_BACK    = 9 //订单退回
)

// 通知类型
var CONF_NOTIFY_TYPE = map[int]string{
	SEND_FAILURE:     "SEND_FAILURE",     //物流下单失败
	SEND_SUCCESS:     "SEND_SUCCESS",     //物流下单成功
	COLLECTING:       "COLLECTING",       //揽收中
	COLLECTED:        "COLLECTED",        //已揽收
	TRANSIT:          "TRANSIT",          //运输中
	SIGNED:           "SIGNED",           //已签收
	CANCELED:         "CANCELED",         //已取消
	CANCELED_INVALID: "CANCELED_INVALID", //订单作废
	CANCELED_BACK:    "CANCELED_BACK",    //订单退回
}

// 通知类型
var CONF_NOTIFY_CATEGORY = map[int]string{
	NOTIFY_CATEGORY_SENT:              "ExpressOrderSent",
	NOTIFY_CATEGORY_RECEIPTING:        "ExpressOrderReceipting",
	NOTIFY_CATEGORY_GOT:               "ExpressOrderGot",
	NOTIFY_CATEGORY_SIGNEDFROMEXPRESS: "ExpressOrderSignedFromExpress",
	NOTIFY_CATEGORY_CANCELLED:         "ExpressOrderCancelled",
	NOTIFY_CATEGORY_AIGUO_CANCEL:      "BusinessOrderCancelled",
}
