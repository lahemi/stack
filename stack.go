// Licensed CC0, so go nuts.
package stack

import "errors"

var emptyStackError = errors.New("Attempting to Pop an empty stack")

type Stack struct {
	s []interface{}
}

func (s *Stack) Push(a interface{}) {
	s.s = append(s.s, a)
}

func (s *Stack) PopE() (interface{}, error) {
	if len(s.s) == 0 {
		return nil, emptyStackError
	}
	last := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return last, nil
}

func (s *Stack) Pop() interface{} {
	rv, err := s.PopE()
	if err != nil {
		return nil
	}
	return rv
}

func (s *Stack) Dup() error {
	rv, err := s.PopE()
	if err != nil {
		return err
	}
	s.Push(rv)
	s.Push(rv)
	return nil
}

func (s *Stack) Drop() {
	s.Pop()
}

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
