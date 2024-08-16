package aiguoclient_v2

const (
	BASE_URL_PRODUCT = "http://aiguo-express.aiguo.tech/"
	BASE_URL_DEV     = "http://testing-aiguo-express.aiguo.tech/"
)

type Client struct {
	AppID   int
	AppKey  string
	BaseURL string
}

func NewClient(appID int, appKey string, baseURL string) *Client {
	return &Client{
		AppID:   appID,
		AppKey:  appKey,
		BaseURL: baseURL,
	}
}
