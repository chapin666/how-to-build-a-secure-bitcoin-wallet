package engine

// OrderBook - 交易委托账本
type OrderBook struct {
	BuyOrders  []Order
	SellOrders []Order
}

// addBuyOrder - Add a buy order to the order book
// bid 方委托单是按升序排列，序号最大的委托单价最高
func (book *OrderBook) addBuyOrder(order Order) {
	n := len(book.BuyOrders)
	var i int

	// 查找比order.price小的数据
	for i = n - 1; i >= 0; i-- {
		buyOrder := book.BuyOrders[i]
		if buyOrder.Price < order.Price {
			break
		}
	}

	// 根据price排序（升序）入队列
	if i == n-1 {
		book.BuyOrders = append(book.BuyOrders, order)
	} else {
		copy(book.BuyOrders[i+1:], book.BuyOrders[i:])
		book.BuyOrders[i] = order
	}
}

// addSellOrder - add a sell order to the order book
// ask 方委托单是按降序排列，序号最大的委托单价最低
func (book *OrderBook) addSellOrder(order Order) {
	n := len(book.SellOrders)
	var i int

	// 查询比order.price大的数据
	for i = n - 1; i >= 0; i-- {
		sellOrder := book.SellOrders[i]
		if sellOrder.Price > order.Price {
			break
		}
	}

	// 降序排序
	if i == n-1 {
		book.SellOrders = append(book.SellOrders, order)
	} else {
		copy(book.SellOrders[i+1:], book.SellOrders[i:])
		book.SellOrders[i] = order
	}

}

// removeBuyOrder - remove a buy order from the order book at given index
func (book *OrderBook) removeBuyOrder(idx int) {
	book.BuyOrders = append(book.BuyOrders[:idx], book.BuyOrders[idx+1:]...)
}

// removeSellOrder - remove a sell order from the order book at given index
func (book *OrderBook) removeSellOrder(idx int) {
	book.SellOrders = append(book.SellOrders[:idx], book.SellOrders[idx+1:]...)
}
