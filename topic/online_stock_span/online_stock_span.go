package online_stock_span

// StockSpanner object will be instantiated and called as such:
// obj := Constructor();
// param_1 := obj.Next(price);

type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (spanner *StockSpanner) Next(price int) int {
	if len(spanner.stack) == 0 || spanner.stack[len(spanner.stack)-1][0] > price {
		spanner.stack = append(spanner.stack, [2]int{price, 1})
		return 1
	}

	count := 1

	for len(spanner.stack) != 0 && spanner.stack[len(spanner.stack)-1][0] <= price {
		count += spanner.stack[len(spanner.stack)-1][1]
		spanner.stack = spanner.stack[:len(spanner.stack)-1]
	}

	spanner.stack = append(spanner.stack, [2]int{price, count})

	return count
}
