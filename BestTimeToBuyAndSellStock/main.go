package main

func main() {}

func maxProfit(prices []int) int {
	boughtPrice := prices[0] // 初日に買ったと仮定するところからスタート
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		todayPrice := prices[i]

		// NOTE: 売買が同じ日に発生することはない
		if boughtPrice > todayPrice {
			// 今日買えばよりお得になるなら今日買う
			boughtPrice = todayPrice
		} else if todayPrice-boughtPrice > maxProfit {
			// 今日売れば利益が最大になるなら、売って、得た利益を暫定最大利益とする
			maxProfit = todayPrice - boughtPrice
		}
	}
	return maxProfit
}
