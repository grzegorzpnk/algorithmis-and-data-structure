package slidingWindows

// First intuitive approach
//func maxProfit(prices []int) int {
//
//	if len(prices) == 1{
//		return 0
//	}
//
//	if len(prices) == 2 && prices[1] > prices[0] {
//		return prices[1] - prices[0]
//	}
//
//	max_profit := 0
//	for i := 0; i <= len(prices)-2; i++{
//		for x := i+1 ; x <= len(prices)-1; x++{
//			profit := prices[x]-prices[i]
//			if profit > 0 && profit > max_profit{
//				max_profit = profit
//			}
//		}
//	}
//	return max_profit
//
//}

func maxProfit(prices []int) int {

	min_price := prices[0]
	max_profit := 0

	for i := range prices {
		if prices[i] < min_price {
			min_price = prices[i]
		}

		if prices[i]-min_price > max_profit {
			max_profit = prices[i] - min_price
		}
	}

	return max_profit

}
