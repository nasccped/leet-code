package main

import "fmt"

// Struct type to mimic stack features using go lang slices.
type Stack[T any] struct {
	stk []T
}

// Generates a new Stack pointer.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{[]T{}}
}

// Push a new item to a stack pointer.
func (s *Stack[T]) Push(item T) {
	s.stk = append(s.stk, item)
}

// Pop the top item of a stack pointer.
func (s *Stack[T]) Pop() (T, error) {
	var (
		top T
		err error = nil
	)
	if s.Len() == 0 {
		err = fmt.Errorf("the stack is empty")
	} else {
		top = s.stk[s.Len()-1]
		s.stk = s.stk[:s.Len()-1]
	}
	return top, err
}

// Returns the stack length.
func (s *Stack[T]) Len() int {
	return len(s.stk)
}

// Peek the top item of a stack.
func (s *Stack[T]) Peek() *T {
	var pointer *T = nil
	if l := s.Len(); l != 0 {
		pointer = &s.stk[l-1]
	}
	return pointer
}

func longestValidParentheses(s string) int {
	result := 0
	parenStack := NewStack[int]()
	parenStack.Push(-1)
	for i, b := range s {
		switch b {
		case '(':
			parenStack.Push(i)
		default:
			parenStack.Pop()
			if parenStack.Len() == 0 {
				parenStack.Push(i)
			} else {
				result = max(result, i-*parenStack.Peek())
			}
		}
	}
	return result
}

func main() {
	inp := longestValidParentheses("(()")
	fmt.Println(inp)
}
