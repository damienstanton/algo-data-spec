package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(solve(str))
}

// handroll a stack
type bytestack []byte

func (s bytestack) pop() byte {
	x := s[len(s)-1]
	s = s[:len(s)-1]
	return x
}

func (s bytestack) contains(b byte) bool {
	for _, s := range s {
		if s == b {
			return true
		}
	}
	return false
}

func solve(str string) string {
	opens := bytestack("{[(")
	closes := bytestack("}])")
	stack := make(bytestack, 0)
	for i, c := range []byte(str) {
		if opens.contains(c) {
			stack = append(stack, c)
		}
		if len(stack) == 0 {
			return fmt.Sprintf("%d", i)
		}
		if closes.contains(c) {
			if string(c) == "}" && string(stack.pop()) != "{" {
				return fmt.Sprintf("%d", i)
			}
			if string(c) == "]" && string(stack.pop()) != "[" {
				return fmt.Sprintf("%d", i)
			}
			if string(c) == ")" && string(stack.pop()) != "(" {
				return fmt.Sprintf("%d", i)
			}
			stack.pop()
		}
	}

	return "Success"
}
