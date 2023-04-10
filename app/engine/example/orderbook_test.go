package example

import (
	"github.com/shopspring/decimal"
	"strconv"
	"testing"
	"time"
)

func TestOrder_Add(t *testing.T) {
	orderbook := NewOrderBook("BTC_USDT")
	for i := 0; i < 10000; i++ {
		tm := time.Now().UnixNano()
		id := strconv.FormatInt(tm, 10)
		ob := &Order{
			OrderId:   id,
			Side:      "buy",
			Price:     decimal.NewFromInt(int64(i)),
			TimeStamp: time.Now().Unix(),
			Quantity:  decimal.NewFromInt(10),
		}
		ok := &OrderKey{
			Price:     ob.Price,
			TimeStamp: ob.TimeStamp,
		}
		//os := &Order{
		//	OrderId:   id,
		//	Side:      "buy",
		//	Price:     decimal.NewFromInt(int64(i)),
		//	TimeStamp: time.Now().Unix(),
		//	Quantity:  decimal.NewFromInt(10),
		//}
		orderbook.Add(ok, ob)
		//orderbook.Add(os)
	}
	it := orderbook.bids.Iterator()
	count := 0
	for it.Next() {
		count++
		k := it.Key()
		t.Log(k)
	}
}

func TestOrder_AddWithSamePrice(t *testing.T) {
	orderbook := NewOrderBook("BTC_USDT")
	// 价格相同，时间不同
	ob := &Order{
		Symbol:    "BTC_USDT",
		OrderId:   "1111110",
		Side:      "buy",
		Price:     decimal.NewFromInt(12),
		TimeStamp: time.Now().UnixNano(),
		Quantity:  decimal.NewFromInt(10),
	}
	ob2 := &Order{
		Symbol:    "BTC_USDT",
		OrderId:   "1111111",
		Side:      "buy",
		Price:     decimal.NewFromInt(12),
		TimeStamp: time.Now().UnixNano(),
		Quantity:  decimal.NewFromInt(10),
	}
	ok := &OrderKey{
		Price:     ob.Price,
		TimeStamp: ob.TimeStamp,
	}
	ok2 := &OrderKey{
		Price:     ob2.Price,
		TimeStamp: ob2.TimeStamp,
	}
	orderbook.Add(ok, ob)
	orderbook.Add(ok2, ob2)
	it := orderbook.bids.Iterator()
	count := 0
	for it.Next() {
		count++
		k := it.Key()
		t.Log(k)
	}
	// &{1111110  BTC_USDT 12 10 buy  1681119356654105000}
	// &{1111111  BTC_USDT 12 10 buy  1681119356654120000}
}

func TestOrder_AddWithDiffPrice(t *testing.T) {
	orderbook := NewOrderBook("BTC_USDT")
	tm := time.Now().UnixNano()
	// 价格不同，时间相同
	ob := &Order{
		Symbol:    "BTC_USDT",
		OrderId:   "1111110",
		Side:      "buy",
		Price:     decimal.NewFromInt(8),
		TimeStamp: tm,
		Quantity:  decimal.NewFromInt(10),
	}
	ob2 := &Order{
		Symbol:    "BTC_USDT",
		OrderId:   "1111111",
		Side:      "buy",
		Price:     decimal.NewFromInt(12),
		TimeStamp: tm,
		Quantity:  decimal.NewFromInt(10),
	}
	ok := &OrderKey{
		Price:     ob.Price,
		TimeStamp: ob.TimeStamp,
	}
	ok2 := &OrderKey{
		Price:     ob2.Price,
		TimeStamp: ob2.TimeStamp,
	}
	orderbook.Add(ok, ob)
	orderbook.Add(ok2, ob2)
	it := orderbook.bids.Iterator()
	count := 0
	for it.Next() {
		count++
		k := it.Key()
		od := it.Value().(*Order)
		od.Quantity = decimal.NewFromInt(9999)
		t.Log(k)
		t.Log(od)
	}
}
