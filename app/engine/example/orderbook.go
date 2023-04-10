package example

import "github.com/emirpasic/gods/trees/redblacktree"

type OrderBook struct {
	symbol     string
	bids       *redblacktree.Tree // 买单列表
	asks       *redblacktree.Tree // 卖单列表
	matchChan  chan struct{}      // order channel 异步顺序处理订单
	chanAdd    chan Order         // 添加订单 channel
	chanCancel chan string        // 移除订单 channel
}

func NewOrderBook(symbol string) *OrderBook {
	return &OrderBook{
		symbol:     symbol,
		bids:       redblacktree.NewWith(OrderBidComparator),
		asks:       redblacktree.NewWith(OrderAskComparator),
		chanAdd:    make(chan Order, 1000000),
		chanCancel: make(chan string, 1000000),
	}
}

func (o *OrderBook) Add(ok *OrderKey, order *Order) {
	switch order.Side {
	case "buy":
		o.bids.Put(ok, order)
	case "sell":
		o.asks.Put(ok, order)
	default:
		o.bids.Put(ok, order)
	}
}

func (o *OrderBook) Remove(orderId string) {

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
			return -1
		} else if orderA.TimeStamp > orderB.TimeStamp {
			return 1
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
