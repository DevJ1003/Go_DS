package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {

	if len(s.items) == 0 {
		return -1
	}

	item := s.items[len(s.items)-1]
	items := s.items[0 : len(s.items)-1]
	s.items = items

	return item
}

func main() {
	s := Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
