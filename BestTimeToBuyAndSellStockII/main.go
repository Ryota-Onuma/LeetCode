package main

func main() {}

func maxProfit(prices []int) int {
	var profit int
	boughtPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		todayPrice := prices[i]
		// 最終日かつ利益が出るなら利確しておしまい
		isFinalDay := i == len(prices)-1
		if isFinalDay {
			if boughtPrice < todayPrice {
				profit += todayPrice - boughtPrice
			}
			break
		}

		// 今日をピークに明日値下がりするなら今日利確して、買い直す
		// you can buy it then immediately sell it on the same day.なのでOK
		tomorrowPrice := prices[i+1]
		if boughtPrice < todayPrice && todayPrice > tomorrowPrice {
			profit += todayPrice - boughtPrice
			boughtPrice = todayPrice
		}

		// 買ったときより今日のほうが安いなら、今日買ったことにしてしまう
		if boughtPrice > todayPrice {
			boughtPrice = todayPrice
		}
	}
	return profit
}
