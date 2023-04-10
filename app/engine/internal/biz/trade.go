package biz

import "github.com/shopspring/decimal"

type Trade struct {
	Symbol      string // 交易对，eg:BTC_USDT
	TradeId     string // 成交id
	TakerId     string // 买单id
	TakerUserId string
	MakerUserId string
	MakerId     string          // 卖单id
	Price       decimal.Decimal // 成交价格
	Quantity    decimal.Decimal // 成交数量
	TimeStamp   int64           // 成交时间
}
