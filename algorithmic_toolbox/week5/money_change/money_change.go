package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	coins := []int{1, 3, 4}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("error converting input to int: %v", err)
			os.Exit(1)
		}
		fmt.Println(change(n, coins))
	}
}

func change(m int, coins []int) int {
	dp := make([]int, m+1)
	dp[0] = 0

	for i := 1; i <= m; i++ {
		dp[i] = math.MaxInt64
		for _, coin := range coins {
			if i >= coin {
				res := dp[i-coin]
				if res < dp[i] && res+1 < dp[i] {
					new := res + 1
					dp[i] = new
				}
			}
		}
	}

	return dp[m]
}
