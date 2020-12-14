package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := make([]string, 0, 2)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(solve(lines[0], lines[1]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(a, b string) int {
	d := make([][]int, len(a)+1)
	m := len(a)
	n := len(b)
	for i := range d {
		d[i] = make([]int, n+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if a[i-1] == b[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				sub := d[i-1][j]
				if d[i][j-1] < sub {
					sub = d[i][j-1]
				}
				if d[i-1][j-1] < sub {
					sub = d[i-1][j-1]
				}
				d[i][j] = sub + 1
			}
		}

	}
	return d[m][n]
}
