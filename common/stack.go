package common

import "errors"

var errPopFromEmptyStack = errors.New("unable to pop from empty stack")
var errPeekFromEmptyStack = errors.New("unable to peek from empty stack")

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(el T) {
	s.stack = append(s.stack, el)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.stack) == 0 {
		var defaultValue T
		return defaultValue, errPopFromEmptyStack
	}
	result := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return result, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.stack) == 0 {
		var defaultValue T
		return defaultValue, errPeekFromEmptyStack
	}
	return s.stack[len(s.stack)-1], nil
}
