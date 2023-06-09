package biz

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type Engine struct {
	orderBooks    map[string]*OrderBook //币对到订单簿的映射
	orderBookLock sync.RWMutex
	ctx           context.Context
}

func NewEngine(symbols []string) (*Engine, error) {

	e := &Engine{
		orderBooks: make(map[string]*OrderBook),
	}
	ctx, _ := context.WithCancel(context.TODO())
	for _, symbol := range symbols {
		orderBook := NewOrderBook(symbol)
		go orderBook.Match()
		orderBook.ctx = ctx
		e.orderBooks[symbol] = orderBook
	}
	return e, nil
}

func (e *Engine) Open(symbol string) error {
	e.orderBookLock.Lock()
	defer e.orderBookLock.Unlock()
	if _, ok := e.orderBooks[symbol]; ok {
		return errors.New("started")
	}
	e.orderBooks[symbol] = NewOrderBook(symbol)
	return nil
}

func (e *Engine) Close(symbol string) error {
	e.orderBookLock.Lock()
	defer e.orderBookLock.Unlock()
	if _, ok := e.orderBooks[symbol]; !ok {
		return errors.New("not find")
	}
	delete(e.orderBooks, symbol)
	// todo 完善逻辑
	return nil
}

func (e *Engine) Add(order *Order) error {
	if order.Quantity.Sign() <= 0 || order.Price.Sign() <= 0 {
		return fmt.Errorf("invalid order quantity or price")
	}
	if _, ok := e.orderBooks[order.Symbol]; !ok {
		return errors.New("not find")
	}
	return e.orderBooks[order.Symbol].AddOrder(order)
}

func (e *Engine) Cancel(symbol, orderId string) error {
	e.orderBookLock.RLock()
	if _, ok := e.orderBooks[symbol]; !ok {
		return errors.New("not find")
	}
	return e.orderBooks[symbol].CancelOrder(orderId)
}
