package onezero

// 错误resp
type CommonErrorResponse struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Status  int    `json:"status"` //The lifetime of the token in seconds
	TraceId string `json:"traceId"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"` //The lifetime of the token in seconds
}

// ----------持仓情况返回-------------------------
type MarginAccountPositionResponse struct {
	Total int                         `json:"total"`
	Data  []MarginAccountPositionItem `json:"data"`
}

type MarginAccountPositionItem struct {
	//must
	Position        string               `json:"position"` //Net position where positive is BUY and negative is SELL
	CoreSymbol      string               `json:"coreSymbol"`
	TierStates      []TierStatesItem     `json:"tierStates"`
	VwapPrecision   int                  `json:"vwapPrecision"`   //The number of decimal places that VWAP is calculated at
	EntityPositions []EntityPositionItem `json:"entityPositions"` //Entity net positions that make up this core symbol net position
	//option
	Vwap                           string `json:"vwap,omitempty"` //VWAP of the net position. Only set if the net position is non-zero
	FifoVWAP                       string `json:"fifoVWAP,omitempty"`
	UnrealizedPnLInAccountCurrency string `json:"unrealizedPnLInAccountCurrency,omitempty"`
	UnrealizedPnL                  struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"unrealizedPnL,omitempty"`
	MarginRate                   string `json:"marginRate,omitempty"`
	Margin                       string `json:"margin,omitempty"`
	FloatingPnLInAccountCurrency string `json:"floatingPnLInAccountCurrency,omitempty"`
	FloatingPnL                  struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"floatingPnL,omitempty"`
	PositionInAccountCurrency string `json:"positionInAccountCurrency,omitempty"`
}

type TierStatesItem struct {
	MarginRate                string `json:"marginRate"`
	PositionInAccountCurrency string `json:"positionInAccountCurrency"`
	MarginInAccountCurrency   string `json:"marginInAccountCurrency"`
}

type EntityPositionItem struct {
	//must
	Position          string `json:"position"`
	EntityId          string `json:"entityId"`
	EntitySymbol      string `json:"entitySymbol"`
	EntityDisplayName string `json:"entityDisplayName"`
	SettlementTimeUTC string `json:"settlementTimeUTC"`
	//option
	Vwap                           string `json:"vwap"`
	FifoVWAP                       string `json:"fifoVWAP"`
	UnrealizedPnLInAccountCurrency string `json:"unrealizedPnLInAccountCurrency"`
	MarketPrice                    string `json:"marketPrice"`
	UnrealizedPnL                  struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"unrealizedPnL"`
}
