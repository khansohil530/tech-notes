package main

func maxProfit(prices []int) int {
	buy := prices[0]
	var profit int
	for _, price := range prices {
		buy = min(price, buy)
		profit = max(profit, price-buy)
	}
	return profit
}
