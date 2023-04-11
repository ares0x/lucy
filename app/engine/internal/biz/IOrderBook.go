package biz

// IOrderBook orderbook 需要实现的方法
type IOrderBook interface {
	AddOrder(order *Order) error
	CancelOrder(orderId string) error
	Match()
}

// 除了使用红黑树，我们还可以选择其他的数据结构来实现订单簿的主体结构，包含二维链表，跳表 等等，只要他们实现上述三个方法
// 关于数据结构的选择，需要考虑的点有一下几个：
// 1.需要一个可以横向和纵向都可以扩展的结构：如，横向保存不同的价格，纵向则用来存储价格相同的订单
// 2.需要比较好的性能（从这个角度看，链表不是一个好的选择，在查找和删除某个节点的值是需要遍历）
