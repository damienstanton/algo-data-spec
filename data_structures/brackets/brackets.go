package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(solve(str))
}

// handroll a stack struct with the expected methods
type runestack struct {
	elements []rune
	pos      int
}

func newStackFrom(bb []rune) *runestack {
	return &runestack{bb, 0}
}

func (s *runestack) push(b rune, i int) {
	s.elements = append(s.elements, b)
	s.pos = i
}

func (s *runestack) pop() rune {
	n := len(s.elements)
	x := s.elements[n-1]
	s.elements = s.elements[:n-1]
	return x
}

func (s *runestack) contains(b rune) bool {
	for _, e := range s.elements {
		if b == e {
			return true
		}
	}
	return false
}

func pair(c rune) rune {
	switch c {
	case '}':
		return '{'
	case ']':
		return '['
	case ')':
		return '('
	default:
		return rune(0)
	}
}

// use the stack to balance parentheses
func solve(str string) string {
	if len(str) <= 1 {
		return "1"
	}
	opens := newStackFrom([]rune("{[("))
	closes := newStackFrom([]rune("}])"))
	stack := newStackFrom([]rune{})

	for i, c := range str {
		if opens.contains(c) {
			stack.push(c, i)
		} else {
			if closes.contains(c) {
				if len(stack.elements) == 0 {
					return fmt.Sprintf("%d", i+1)
				}
				want := pair(c)
				top := stack.pop()
				if top != want {
					return fmt.Sprintf("%d", i+1)
				}
			}
		}
	}

	if len(stack.elements) == 0 {
		return "Success"
	}

	return fmt.Sprintf("%d", stack.pos)
}
