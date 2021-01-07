package main

import "fmt"

func main() {
	// line 1
	var n int
	fmt.Scan(&n)

	// line2
	vs := make([]int, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		vs[i] = v
	}
}
