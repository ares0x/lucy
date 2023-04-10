package biz

import (
	"context"
	errors2 "errors"
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/shopspring/decimal"
	"lucy/app/engine/internal/constvar"
	"time"
)

type OrderBook struct {
	symbol     string
	bids       *redblacktree.Tree // 买单列表
	asks       *redblacktree.Tree // 卖单列表
	matchChan  chan struct{}      // order channel 异步顺序处理订单
	chanAdd    chan Order         // 添加订单 channel
	chanCancel chan string        // 移除订单 channel
	trades     []*Trade           // 成交列表
	ctx        context.Context
}

func NewOrderBook(symbol string) *OrderBook {
	return &OrderBook{
		symbol:     symbol,
		bids:       redblacktree.NewWith(OrderBidComparator),
		asks:       redblacktree.NewWith(OrderAskComparator),
		chanAdd:    make(chan Order, 1000000),
		chanCancel: make(chan string, 1000000),
		matchChan:  make(chan struct{}),
	}
}

// AddOrder 下单
func (ob *OrderBook) AddOrder(order *Order) error {
	select {
	case ob.chanAdd <- *order: // 订单加入有缓冲的 channel
		return nil
	case <-ob.ctx.Done():
		return errors2.New("closed")
	}
}

// CancelOrder 撤单
func (ob *OrderBook) CancelOrder(orderId string) error {
	select {
	case ob.chanCancel <- orderId:
		return nil
	case <-ob.ctx.Done():
		return errors2.New("closed")
	}
}

// Match 开启订单簿的挂单和撤单逻辑
func (ob *OrderBook) Match() {
	for {
		select {
		case order := <-ob.chanAdd:
			ob.add(&order)
		case orderId := <-ob.chanCancel:
			ob.cancel(orderId)
		case <-ob.ctx.Done():
			return
		}
	}
}

func (ob *OrderBook) add(order *Order) error {
	switch order.Side {
	case string(constvar.Buy):
		return ob.addBid(order)
	case string(constvar.Sell):
		return ob.addAsk(order)
	}
	return errors2.New("not found side")
}

func (ob *OrderBook) cancel(orderId string) error {
	// getByOrderId

	return nil
}

func (ob *OrderBook) addBid(order *Order) error {
	switch order.Type {
	case string(constvar.Limit):
		return ob.addBidLimit(order)
	case string(constvar.Market):
		return ob.addBidMarket(order)
	}
	return errors2.New("not found type")
}

func (ob *OrderBook) addAsk(order *Order) error {
	switch order.Type {
	case string(constvar.Limit):
		return ob.addBidLimit(order)
	case string(constvar.Market):
		return ob.addBidMarket(order)
	}
	return errors2.New("not found type")
}

// addBidLimit 限价 买
func (ob *OrderBook) addBidLimit(order *Order) error {
	trades := make([]Trade, 0)
	treeIterator := ob.asks.Iterator()
	for treeIterator.Next() && !order.Quantity.IsNegative() {
		askOrder := treeIterator.Value().(*Order)
		if order.Price.GreaterThanOrEqual(askOrder.Price) { // 买单价格大于等于最佳匹配价格
			if askOrder.Quantity.GreaterThanOrEqual(order.Quantity) {
				trade := Trade{
					Symbol:      order.Symbol,
					TradeId:     "",
					TakerId:     order.OrderId,
					MakerId:     askOrder.OrderId,
					TakerUserId: order.UserId,
					MakerUserId: askOrder.UserId,
					Price:       askOrder.Price,
					Quantity:    order.Quantity,
					TimeStamp:   time.Now().UnixNano(),
				}
				trades = append(trades, trade)
				remainingQuantity := askOrder.Quantity.Sub(order.Quantity) // 剩余数量
				if remainingQuantity.IsPositive() {                        // 剩余资产数量大于0
					askOrder.Quantity = remainingQuantity
				} else { // 剩余资产小于等于 0，直接更新
					orderKey := &OrderKey{
						Price:     askOrder.Price,
						TimeStamp: askOrder.TimeStamp,
						OrderId:   askOrder.OrderId,
					}
					ob.cancelAsk(orderKey)
				}
			} else {
				trade := Trade{
					Symbol:      order.Symbol,
					TradeId:     "",
					TakerId:     order.OrderId,
					MakerId:     askOrder.OrderId,
					TakerUserId: order.UserId,
					MakerUserId: askOrder.UserId,
					Price:       askOrder.Price,
					Quantity:    order.Quantity,
					TimeStamp:   time.Now().UnixNano(),
				}
				trades = append(trades, trade)
				order.Quantity = order.Quantity.Sub(askOrder.Quantity)
				orderKey := &OrderKey{
					Price:     askOrder.Price,
					TimeStamp: askOrder.TimeStamp,
					OrderId:   askOrder.OrderId,
				}
				ob.cancelAsk(orderKey)
			}
		} else {
			if order.Quantity.IsPositive() {
				orderKey := &OrderKey{
					OrderId:   order.OrderId,
					Price:     order.Price,
					TimeStamp: order.TimeStamp,
				}
				ob.bids.Put(orderKey, order) //
			}
		}

	}
	if len(trades) > 0 {
		// TODO 推送成交
	}
	return nil
}

// addBidLimit 市价
func (ob *OrderBook) addBidMarket(order *Order) error {
	trades := make([]Trade, 0)
	treeIterator := ob.asks.Iterator()
	for treeIterator.Next() && !order.Quantity.IsNegative() {
		askOrder := treeIterator.Value().(*Order)
		if askOrder.Quantity.GreaterThanOrEqual(order.Quantity) {
			trade := Trade{
				Symbol:      order.Symbol,
				TradeId:     "",
				TakerId:     order.OrderId,
				MakerId:     askOrder.OrderId,
				TakerUserId: order.UserId,
				MakerUserId: askOrder.UserId,
				Price:       askOrder.Price,
				Quantity:    order.Quantity,
				TimeStamp:   time.Now().UnixNano(),
			}
			trades = append(trades, trade)
			remainingQuantity := askOrder.Quantity.Sub(order.Quantity) // 剩余数量
			if remainingQuantity.IsPositive() {                        // 剩余资产数量大于0
				askOrder.Quantity = remainingQuantity
			} else { // 剩余资产小于等于 0，直接更新
				orderKey := &OrderKey{
					Price:     askOrder.Price,
					TimeStamp: askOrder.TimeStamp,
					OrderId:   askOrder.OrderId,
				}
				ob.cancelAsk(orderKey)
			}
		} else {
			trade := Trade{
				Symbol:      order.Symbol,
				TradeId:     "",
				TakerId:     order.OrderId,
				MakerId:     askOrder.OrderId,
				TakerUserId: order.UserId,
				MakerUserId: askOrder.UserId,
				Price:       askOrder.Price,
				Quantity:    order.Quantity,
				TimeStamp:   time.Now().UnixNano(),
			}
			trades = append(trades, trade)
			order.Quantity = order.Quantity.Sub(askOrder.Quantity)
			orderKey := &OrderKey{
				Price:     askOrder.Price,
				TimeStamp: askOrder.TimeStamp,
				OrderId:   askOrder.OrderId,
			}
			ob.cancelAsk(orderKey)
		}
	}
	// 判断是否完全成交，如果没有则撤单
	if order.Quantity.IsZero() {
		//TODO 推送成交单
	}
	//TODO 撤销剩余委托单数量
	return nil
}

// addAskLimit 限价 卖
func (ob *OrderBook) addAskLimit(order *Order) error {
	trades := make([]Trade, 0)
	treeIterator := ob.bids.Iterator()
	for treeIterator.Next() && !order.Quantity.IsNegative() {
		bidOrder := treeIterator.Value().(*Order)
		if order.Price.LessThanOrEqual(bidOrder.Price) { // 买单价格大于等于最佳匹配价格
			if bidOrder.Quantity.GreaterThanOrEqual(order.Quantity) {
				trade := Trade{
					Symbol:      order.Symbol,
					TradeId:     "",
					TakerId:     order.OrderId,
					MakerId:     bidOrder.OrderId,
					TakerUserId: order.UserId,
					MakerUserId: bidOrder.UserId,
					Price:       bidOrder.Price,
					Quantity:    order.Quantity,
					TimeStamp:   time.Now().UnixNano(),
				}
				trades = append(trades, trade)
				remainingQuantity := bidOrder.Quantity.Sub(order.Quantity) // 剩余数量
				if remainingQuantity.IsPositive() {                        // 剩余资产数量大于0
					bidOrder.Quantity = remainingQuantity
				} else { // 剩余资产小于等于 0，直接更新
					orderKey := &OrderKey{
						OrderId:   bidOrder.OrderId,
						Price:     bidOrder.Price,
						TimeStamp: bidOrder.TimeStamp,
					}
					ob.cancelBid(orderKey)
				}
			} else {
				trade := Trade{
					Symbol:      order.Symbol,
					TradeId:     "",
					TakerId:     order.OrderId,
					MakerId:     bidOrder.OrderId,
					TakerUserId: order.UserId,
					MakerUserId: bidOrder.UserId,
					Price:       bidOrder.Price,
					Quantity:    order.Quantity,
					TimeStamp:   time.Now().UnixNano(),
				}
				trades = append(trades, trade)
				order.Quantity = order.Quantity.Sub(bidOrder.Quantity)
				orderKey := &OrderKey{
					OrderId:   bidOrder.OrderId,
					Price:     bidOrder.Price,
					TimeStamp: bidOrder.TimeStamp,
				}
				ob.cancelBid(orderKey)
			}
		} else {
			if order.Quantity.IsPositive() {
				orderKey := &OrderKey{
					OrderId:   order.OrderId,
					Price:     order.Price,
					TimeStamp: order.TimeStamp,
				}
				ob.asks.Put(orderKey, order) //
			}
		}

	}
	if len(trades) > 0 {
		// TODO 推送成交
	}
	return nil
}

// addAskMarket 市价
func (ob *OrderBook) addAskMarket(order *Order) error {
	trades := make([]Trade, 0)
	treeIterator := ob.bids.Iterator()
	for treeIterator.Next() && !order.Quantity.IsNegative() {
		bidOrder := treeIterator.Value().(*Order)
		if bidOrder.Quantity.GreaterThanOrEqual(order.Quantity) {
			trade := Trade{
				Symbol:      order.Symbol,
				TradeId:     "",
				TakerId:     order.OrderId,
				MakerId:     bidOrder.OrderId,
				TakerUserId: order.UserId,
				MakerUserId: bidOrder.UserId,
				Price:       bidOrder.Price,
				Quantity:    order.Quantity,
				TimeStamp:   time.Now().UnixNano(),
			}
			trades = append(trades, trade)
			remainingQuantity := bidOrder.Quantity.Sub(order.Quantity) // 剩余数量
			if remainingQuantity.IsPositive() {                        // 剩余资产数量大于0
				bidOrder.Quantity = remainingQuantity
			} else { // 剩余资产小于等于 0，直接更新
				orderKey := &OrderKey{
					OrderId:   bidOrder.OrderId,
					Price:     bidOrder.Price,
					TimeStamp: bidOrder.TimeStamp,
				}
				ob.cancelBid(orderKey)
			}
		} else {
			trade := Trade{
				Symbol:      order.Symbol,
				TradeId:     "",
				TakerId:     order.OrderId,
				MakerId:     bidOrder.OrderId,
				TakerUserId: order.UserId,
				MakerUserId: bidOrder.UserId,
				Price:       bidOrder.Price,
				Quantity:    order.Quantity,
				TimeStamp:   time.Now().UnixNano(),
			}
			trades = append(trades, trade)
			order.Quantity = order.Quantity.Sub(bidOrder.Quantity)
			orderKey := &OrderKey{
				OrderId:   bidOrder.OrderId,
				Price:     bidOrder.Price,
				TimeStamp: bidOrder.TimeStamp,
			}
			ob.cancelBid(orderKey)
		}
	}
	// 判断是否完全成交，如果没有则撤单
	if order.Quantity.IsZero() {
		//TODO 推送成交单
	}
	//TODO 撤销剩余委托单数量
	return nil
}

func (ob *OrderBook) cancelBid(orderKey *OrderKey) {
	ob.bids.Remove(orderKey)
}

func (ob *OrderBook) cancelAsk(orderKey *OrderKey) {
	ob.asks.Remove(orderKey)
}

func (ob *OrderBook) Get() {

}

type OrderKey struct {
	OrderId   string
	Price     decimal.Decimal
	TimeStamp int64
}

// OrderBidComparator 买单定序规则
func OrderBidComparator(a, b interface{}) int {
	orderA, orderB := a.(*OrderKey), b.(*OrderKey)
	switch {
	case orderA.Price.GreaterThan(orderB.Price):
		return -1
	case orderA.Price.LessThan(orderB.Price):
		return 1
	default: // 价格相同，按时间优先
		if orderA.TimeStamp < orderB.TimeStamp {
			return 1
		} else if orderA.TimeStamp > orderB.TimeStamp {
			return -1
		} else {
			return 0
		}
	}
}

// OrderAskComparator 卖单定序规则
func OrderAskComparator(a, b interface{}) int {
	orderA, orderB := a.(*OrderKey), b.(*OrderKey)
	switch {
	case orderA.Price.GreaterThan(orderB.Price):
		return 1
	case orderA.Price.LessThan(orderB.Price):
		return -1
	default: // 价格相同，按时间优先
		if orderA.TimeStamp < orderB.TimeStamp {
			return -1
		} else if orderA.TimeStamp > orderB.TimeStamp {
			return 1
		} else {
			return 0
		}
	}
}
