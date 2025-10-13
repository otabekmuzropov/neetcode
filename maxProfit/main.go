package main

import "fmt"

func main() {
	prices := []int{7, 8, 9, 3, 1, 4, 6}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	L := 0
	profit := 0
	for R := 1; R < len(prices); R++ {
		diff := prices[R] - prices[L]
		if diff < 0 {
			L = R
		} else {
			profit = max(profit, prices[R]-prices[L])
		}
	}

	return profit
}
