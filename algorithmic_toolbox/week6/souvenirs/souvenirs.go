package main

import (
	"fmt"
)

func main() {
	var n int
	// line 1
	fmt.Scan(&n)

	// line 2
	sum := 0
	vs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var n int
		fmt.Scan(&n)
		sum += n
		vs[i] = n
	}

	partitions := 3
	segments := sum / partitions
	if sum%partitions != 0 {
		fmt.Println(0)
	} else {
		fmt.Println(solve(segments, vs))
	}
}

func solve(n int, vs []int) int {
	res := 0
	dp := make([][]int, len(vs))
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= len(vs)-1; i++ {
		for w := 1; w <= n; w++ {
			dp[i][w] = dp[i-1][w]
			if vs[i] <= w {
				val := dp[i-1][w-vs[i]] + vs[i]
				if dp[i][w] < val {
					dp[i][w] = val
				}
			}
			// maximal, so increment the number of successful partitions
			if dp[i][w] == n {
				res++
			}
		}
	}

	if res < 3 {
		return 0
	}
	return 1
}
