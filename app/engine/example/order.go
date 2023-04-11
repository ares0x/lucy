package example

import "github.com/shopspring/decimal"

type Order struct {
	OrderId   string          `json:"orderId"`   // 订单id
	UserId    string          `json:"userId"`    // 用户id
	Symbol    string          `json:"symbol"`    // 交易对，eg:BTC_USDT
	Price     decimal.Decimal `json:"price"`     // 价格
	Quantity  decimal.Decimal `json:"quantity"`  // 数量
	Side      string          `json:"side"`      // 订单方向 buy/sell
	Type      string          `json:"type"`      // 订单类型 limit/market
	TimeStamp int64           `json:"timeStamp"` // 创建时间
}

type OrderKey struct {
	OrderId   string // 订单id
	Price     decimal.Decimal
	TimeStamp int64
}
