package main

import "fmt"

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				if dp[i] > 1+dp[i-c] {
					dp[i] = 1 + dp[i-c]
				}
			}
		}
	}
	if dp[amount] != amount+1 {
		return dp[amount]
	} else {
		return -1
	}
}

func main() {
	coins := [4]int{1, 3, 4, 5}
	amount := 97
	fmt.Println(CoinChange(coins[:], amount))
}
