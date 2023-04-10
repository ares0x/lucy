package biz

type OrderBookRepo interface {
	AddOrder(order *Order) error
	CancelOrder(orderId string) error
	Match()
}
