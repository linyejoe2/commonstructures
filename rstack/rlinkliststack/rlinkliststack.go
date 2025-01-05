package rlinkliststack

import (
	"github.com/linyejoe2/commonstructures/rlinkedlist"
)

type RLinkListStack[T any] struct {
	_list *rlinkedlist.ListNode[T]
}

func NewRLinkListStack[T any]() *RLinkListStack[T] {
	return &RLinkListStack[T]{
		_list: nil,
	}
}

func (s *RLinkListStack[T]) Push(val T) {
	newHead := rlinkedlist.NewListNode(val)
	newHead.Next = s._list
	s._list = newHead
}

func (s *RLinkListStack[T]) Pop() (val T) {
	if s._list == nil {
		return
	}

	val = s._list.Val
	s._list = s._list.Next
	return
}

func (s *RLinkListStack[T]) Peek() (val T) {
	if s._list == nil {
		return
	}

	val = s._list.Val
	return
} 

func (s *RLinkListStack[T]) Size() (int) {
	:
}

func (s *RLinkListStack[T]) PopAndPrintAll() {
	print(s.Pop(), " ")

	if s._list == nil {
		println("")
		return
	}
	s.PopAndPrintAll()
}
