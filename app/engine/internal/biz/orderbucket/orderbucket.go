package orderbucket

import (
	"github.com/shopspring/decimal"
	"lucy/app/engine/internal/biz"
)

type OrderBucket struct {
	side        string               // 方向：buy/sell
	volume      decimal.Decimal      // 总量
	level       int                  // 层级
	bucketEntry map[string]biz.Order // 桶中委托单 key:orderId value:order
}

// Add 添加委托单
func (o *OrderBucket) Add(order biz.Order) error { // 添加

	return nil
}

// Remove 移除委托单
func (o *OrderBucket) Remove(orderId string) error { // 删除
	return nil
}

// Match 成交
func (o *OrderBucket) Match() {

}

// GetPrice 获取价格
func (o *OrderBucket) GetPrice() decimal.Decimal {

	return decimal.Decimal{}
}

// SetPrice 设置价格
func (o *OrderBucket) SetPrice() {

}

// GetTotalVolume 获取总量 可以修改，如 Binance 显示的是成交总量， OKX 显示的是委托总量）
func (o *OrderBucket) GetTotalVolume() decimal.Decimal {

	return decimal.Decimal{}
}
