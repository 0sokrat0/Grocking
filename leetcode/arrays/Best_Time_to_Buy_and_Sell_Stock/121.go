package besttimetobuyandsellstock

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			profit = prices[i] - minPrice
		}
	}
	fmt.Print(profit)
	return profit
}
