package rlinkliststack

import (
	"github.com/linyejoe2/commonstructures/rlinkedlist"
)

type RLinkListStack[T any] struct {
	_list   *rlinkedlist.ListNode[T]
	_length int
}

func NewRLinkListStack[T any]() *RLinkListStack[T] {
	return &RLinkListStack[T]{
		_list:   nil,
		_length: 0,
	}
}

func (s *RLinkListStack[T]) Push(val T) {
	newHead := rlinkedlist.NewListNode(val)
	newHead.Next = s._list
	s._list = newHead
	s._length++
}

func (s *RLinkListStack[T]) Pop() (val T) {
	if s._list == nil {
		return
	}

	val = s._list.Val
	s._list = s._list.Next
	s._length--
	return
}

func (s *RLinkListStack[T]) Peek() (val T) {
	if s._list == nil {
		return
	}

	val = s._list.Val
	return
}

func (s *RLinkListStack[T]) Size() int {
	return s._length
}

func (s *RLinkListStack[T]) PopAndPrintAll() {
	print(s.Pop(), " ")

	if s._list == nil {
		println("")
		return
	}
	s.PopAndPrintAll()
}

func (s *RLinkListStack[T]) IsEmpty() bool {
	return s._length == 0
}

func (s *RLinkListStack[T]) ToSliceWithNewInstance() (slice *[]T) {
	ele := s._list
	slice = new([]T)
	for ele != nil {
		*slice = append(*slice, ele.Val)
		ele = ele.Next
	}
	return
}

func (s *RLinkListStack[T]) ToSlice() (slice []T) {
	ele := s._list
	for ele != nil {
		slice = append(slice, ele.Val)
		ele = ele.Next
	}
	return
}
