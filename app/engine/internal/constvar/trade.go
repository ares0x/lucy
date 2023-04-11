package constvar

type TradeSide string

const (
	Buy  TradeSide = "buy"  // 买
	Sell TradeSide = "sell" // 卖
)

type TradeType string

const (
	Limit  TradeType = "limit"  // 限价
	Market TradeType = "market" // 市价
)
