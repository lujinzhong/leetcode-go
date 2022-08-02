package main

//给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
//
//在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
//返回 你能获得的 最大 利润 。

func maxProfit(prices []int) int {
	profits := 0
	for k, v := range prices {
		if k == 0 {
			continue
		}
		if v > prices[k-1] { // 后一天比前一天股价高，则卖出
			profits = profits + v - prices[k-1]
		}
	}
	return profits
}

func main() {
}
