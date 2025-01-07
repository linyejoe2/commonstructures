package rlinkedlistqueue

import (
	"github.com/linyejoe2/commonstructures/rlinkedlist"
)

type RLinkListQueue[T any] struct {
	_head   *rlinkedlist.ListNode[T]
	_tail   *rlinkedlist.ListNode[T]
	_length int
}

func NewRLinkListQueue[T any]() *RLinkListQueue[T] {
	return &RLinkListQueue[T]{
		_head: nil,
		_tail: nil,
	}
}

func (q *RLinkListQueue[T]) Enqueue(v T) {
	q._length++

	if q._head == nil {
		q._head = rlinkedlist.NewListNode(v)
		q._tail = q._head
		return
	}
	q._tail.Next = rlinkedlist.NewListNode(v)
	q._tail = q._tail.Next
}

func (q *RLinkListQueue[T]) Dequeue() (v T) {
	if q._head == nil {
		return
	}

	q._length--

	v = q._head.Val
	q._head = q._head.Next
	return
}

func (q *RLinkListQueue[T]) Peek() (v T) {
	if q._head == nil {
		return
	}
	v = q._head.Val
	return
}

func (q *RLinkListQueue[T]) Size() int {
	return q._length
}

func (q *RLinkListQueue[T]) IsEmpty() bool {
	return q._length == 0
}

func (q *RLinkListQueue[T]) ToSlice() (slice []T) {
	return toSliceRecursive(q._head, *new([]T))
}

func toSliceRecursive[T any](head *rlinkedlist.ListNode[T], slice []T) []T {
	if head == nil {
		return slice
	}

	return toSliceRecursive(head.Next, append(slice, head.Val))
}
