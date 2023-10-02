package main

import "fmt"

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(d interface{}) {
	s.data = append(s.data, d)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.check_ifEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	lastVal := s.data[len(s.data)-1]
	s.data = append(s.data[:len(s.data)-1])
	// return lastVal
	return lastVal, nil
}

func (s *Stack) check_ifEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false

}

func (s *Stack) print() {

	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Println(s.data[i])
	}
}

func main() {
	stack := Stack{}
	stack.Push("test")
	stack.Push("test1")
	stack.Pop()
	stack.Push(2)
	stack.print()
	
}
