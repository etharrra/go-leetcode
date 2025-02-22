package solutions

func MaxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		sell := prices[i]
		if sell > buy {
			profit = max(profit, sell-buy)
		} else {
			buy = sell
		}
	}
	// max := -99999
	// for i := 0; i < len(prices); i++ {
	// 	current := prices[i]
	// 	right := prices[i+1:]
	// 	max_right := 0
	// 	if len(right) > 0 {
	// 		max_right = slices.Max(right)
	// 	}
	// 	// fmt.Println(max)
	// 	profit := max_right - current
	// 	if profit > max {
	// 		max = profit
	// 	}
	// }
	// if max <= 0 {
	// 	return 0
	// }
	// return max

	// maxRight := -1
	// min := -1
	// // remove 0
	// // prices = slices.DeleteFunc(prices, func(n int) bool {
	// // 	return n == 0
	// // })
	// for len(prices) > 0 {
	// 	min = slices.Min(prices)
	// 	minIndex := slices.Index(prices, min)
	// 	if len(prices[minIndex+1:]) > 0 {
	// 		// fmt.Println(26)
	// 		maxRight = slices.Max(prices[minIndex+1:])
	// 	}
	// 	if maxRight < 0 {
	// 		// fmt.Println(30)
	// 		prices = slices.Delete(prices, minIndex, minIndex+1)
	// 	} else {
	// 		break
	// 	}
	// }

	// if maxRight < 0 {
	// 	return 0
	// }
	// return maxRight - min

	// min := struct {
	// 	Val   int
	// 	Index int
	// }{prices[len(prices)-1], len(prices) - 1}

	// max := struct {
	// 	Val   int
	// 	Index int
	// }{prices[len(prices)-1], len(prices) - 1}

	// // fmt.Println(min.Val, min.Index)
	// // fmt.Println(max.Val, max.Index)

	// for i := len(prices) - 1; i >= 0; i-- {
	// 	fmt.Println(prices[i])
	// 	if prices[i] < min.Val {
	// 		min.Val = prices[i]
	// 	}
	// 	if prices[i] > max.Val {
	// 		max.Val = prices[i]
	// 	}
	// }
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
