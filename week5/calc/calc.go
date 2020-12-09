package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("error converting input to int: %v", err)
			os.Exit(1)
		}
		minK, seq := solve(n)
		fmt.Println(minK)
		for i := len(seq) - 2; i >= 0; i-- {
			fmt.Printf("%d ", seq[i])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(n int) (int, []int) {
	// first line
	ops := make([]int, n+1)
	for i := 2; i < len(ops); i++ {
		x2 := math.MaxInt64
		x3 := math.MaxInt64
		add := ops[i-1] + 1
		if i%2 == 0 {
			x2 = ops[i/2] + 1
		}
		if i%3 == 0 {
			x3 = ops[i/3] + 1
		}

		ops[i] = min(min(add, x2), x3)
	}
	minK := ops[n]

	// second line
	seq := make([]int, 0, minK+1)
	seq = append(seq, n)
	for i := n; i >= 1; {
		if i%3 == 0 && ops[i]-1 == ops[i/3] {
			seq = append(seq, i/3)
			i /= 3
		} else if i%2 == 0 && ops[i]-1 == ops[i/2] {
			seq = append(seq, i/2)
			i /= 2
		} else {
			seq = append(seq, i-1)
			i--
		}
	}

	return minK, seq
}
