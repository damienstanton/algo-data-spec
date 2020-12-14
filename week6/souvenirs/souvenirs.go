package main

import (
	"fmt"
)

func main() {
	var n int
	// line 1
	fmt.Scan(&n)

	// line 2
	vs := make([]int, n)
	for i := range vs {
		fmt.Scan(&vs[i])
	}

	fmt.Println(solve(n, 3, vs))
}

func solve(n, partitions int, vs []int) int {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	if sum%partitions != 0 {
		return 0
	}
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, (sum/partitions)+1)
	}

	// debug
	for _, row := range dp {
		fmt.Printf("%+v\n", row)
	}

	// TODO
	return 0
}
