// Package stack provides a general stack with some basic stack operation methods.
package stack

import "errors"

var errEmptyStack = errors.New("attempting to Pop an empty stack")

type Stack struct {
	st []interface{}
}

// ( _ → n )
func (s *Stack) Push(a interface{}) {
	s.st = append(s.st, a)
}

// ( n → _ )
func (s *Stack) PopE() (interface{}, error) {
	if len(s.st) == 0 {
		return nil, errEmptyStack
	}
	last := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	return last, nil
}

// ( n → _ )
func (s *Stack) Pop() interface{} {
	rv, err := s.PopE()
	if err != nil {
		return nil
	}
	return rv
}

// ( -- )
func (s *Stack) Depth() int {
	return len(s.st)
}

// ( a → a a )
func (s *Stack) Dup() error {
	rv, err := s.PopE()
	if err != nil {
		return err
	}
	s.Push(rv)
	s.Push(rv)
	return nil
}

// ( a → _ )
func (s *Stack) Drop() {
	s.Pop()
}

// ( a b → b a )
func (s *Stack) Swap() error {
	r2, err := s.PopE()
	r1, err := s.PopE()
	if err != nil {
		return err
	}
	s.Push(r2)
	s.Push(r1)
	return nil
}

// ( a b → a b a )
func (s *Stack) Over() error {
	r2, err := s.PopE()
	r1, err := s.PopE()
	if err != nil {
		return err
	}
	s.Push(r1)
	s.Push(r2)
	s.Push(r1)
	return nil
}

// ( a b c → b c a )
func (s *Stack) Rot() error {
	r3, err := s.PopE()
	r2, err := s.PopE()
	r1, err := s.PopE()
	if err != nil {
		return err
	}
	s.Push(r2)
	s.Push(r3)
	s.Push(r1)
	return nil
}

// Copy n-th element of the stack to the top of it,
// numbering starting from the top of the stack, ie.
// 0 is the top-most element on the stack.
func (s *Stack) Pick(n int) {
	if len(s.st) == 0 || n < 0 {
		// Since the stack can contain different
		// types, we cannot just default to pushing
		// 0 onto the stack like, for instance,
		// pforth does, so we just ignore an empty pick.
		// Also, if the arg is negative, just ignore.
		return
	}
	if n > len(s.st) {
		// As per pforth, if the argument to pick is
		// larger than the number of elements on the
		// stack, we just copy the bottom-most element.
		s.Push(s.st[0])
	}

	// Using len instead of s.Depth `internally`
	// saves us from the function call `overhead`.
	s.Push(s.st[len(s.st)-n-1])
}
