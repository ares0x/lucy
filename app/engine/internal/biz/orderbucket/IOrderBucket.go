package orderbucket

import (
	"github.com/shopspring/decimal"
	"lucy/app/engine/internal/biz"
)

// IOrderBucket orderbucket 是 order 和 orderbook 的中间层，它用来存储相同价格的委托单，即每个 orderbucket 就是一档
type IOrderBucket interface {
	Add(order biz.Order) error       // 添加
	Remove(orderId string) error     // 删除
	Match()                          // 成交
	GetPrice() decimal.Decimal       // 获取价格
	SetPrice()                       // 设置价格
	GetTotalVolume() decimal.Decimal // 获取总量（可以修改，如 Binance 显示的是成交总量， OKX 显示的是委托总量）
}
