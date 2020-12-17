package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(solve(str))
}

type el struct {
	char rune
	pos  int
}

func newElement(c rune, i int) el {
	return el{c, i}
}

type runestack struct {
	elements []el
}

func newStackFrom(s string) runestack {
	elements := make([]el, 0)
	for i, r := range []rune(s) {
		elements = append(elements, el{r, i})
	}
	return runestack{elements}
}

func (s *runestack) push(e el) {
	s.elements = append(s.elements, e)
}

func (s *runestack) pop() el {
	n := len(s.elements) - 1
	res := s.elements[n]
	s.elements = s.elements[:n]
	return res
}

func (s *runestack) isEmpty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}

func (s *runestack) contains(r rune) bool {
	for _, e := range s.elements {
		if e.char == r {
			return true
		}
	}
	return false
}

func pair(r rune) rune {
	switch r {
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

func solve(str string) string {
	stack := runestack{}
	openStack := newStackFrom("{[(")
	closeStack := newStackFrom("}])")

	for i, c := range []rune(str) {
		comp := newElement(c, i)
		if openStack.contains(c) {
			stack.push(newElement(c, i))
		} else if closeStack.contains(c) {
			if stack.isEmpty() {
				return fmt.Sprintf("%d", comp.pos+1)
			}
			top := stack.pop()
			if top.char != pair(c) {
				return fmt.Sprintf("%d", comp.pos+1)
			}
		}
	}

	if !stack.isEmpty() {
		return fmt.Sprintf("%d", stack.elements[0].pos+1)
	}
	return "Success"
}
