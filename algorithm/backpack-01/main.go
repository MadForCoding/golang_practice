package main

import "fmt"

func main() {
	cargoWeight := []int{1, 3, 4}
	cargoValue := []int{15, 20, 30}
	backpackCap := 5
	fmt.Println(backpackMaxValue(cargoWeight, cargoValue, backpackCap))
}

func backpackMaxValue(cargoWeight, cargoValue []int, backpackCap int) int {
	dp := make([][]int, len(cargoWeight))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, backpackCap+1)
	}
	// init
	for j := 0; j < backpackCap; j++ {
		if cargoWeight[0] <= j {
			dp[0][j] = cargoValue[0]
		}
	}
	for i := 1; i < len(cargoWeight); i++ {
		for j := 1; j <= backpackCap; j++ {
			leftWeight := j - cargoWeight[i]
			if leftWeight < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = maxVal(dp[i-1][j], dp[i-1][leftWeight]+cargoValue[i])
			}
		}
	}
	return dp[len(cargoWeight)-1][backpackCap]
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
