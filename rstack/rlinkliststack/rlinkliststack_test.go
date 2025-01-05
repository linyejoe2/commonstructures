package rlinkliststack

import "testing"

func TestMain(t *testing.T) {
	stack := NewRLinkListStack[int]()

	stack.Push(2)
	stack.Push(1)
	stack.Push(3)

	stack.PopAndPrintAll()
}

// func PopAndPrintAll[T any](s *RLinkListStack[T]) {
// 	println(s.Pop())
//
// 	if s._list != nil {
// 		PopAndPrintAll(s)
// 	}
// }
