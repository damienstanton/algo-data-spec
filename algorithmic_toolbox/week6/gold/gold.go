package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	w       int
	n       int
	weights []int
}

func (b *bag) solve() int {
	dp := make([][]int, b.n+1)
	for i := range dp {
		dp[i] = make([]int, b.w+1)
	}

	for i := 1; i <= b.n; i++ {
		for j := 1; j <= b.w; j++ {
			dp[i][j] = dp[i-1][j]
			if b.weights[i] <= j {
				v := dp[i-1][j-b.weights[i]] + b.weights[i]
				if dp[i][j] < v {
					dp[i][j] = v
				}
			}
		}
	}

	return dp[b.n][b.w]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// line 1
	scanner.Scan()
	props := strings.Split(scanner.Text(), " ")
	w, _ := strconv.Atoi(props[0])
	n, _ := strconv.Atoi(props[1])

	// line 2
	scanner.Scan()
	weights := strings.Split(scanner.Text(), " ")
	weights = append([]string{"0"}, weights...)

	bag := &bag{
		w:       w,
		n:       n,
		weights: make([]int, n+1),
	}

	for i := 1; i <= n; i++ {
		n, _ := strconv.Atoi(weights[i])
		bag.weights[i] = n
	}

	fmt.Println(bag.solve())
}
