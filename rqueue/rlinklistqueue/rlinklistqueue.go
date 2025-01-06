package rlinklistqueue

import "github.com/linyejoe2/commonstructures/rlinkedlist"

type RLinkListQueue[T any] struct {
	_list *rlinkedlist.ListNode[T]
}

func NewRLinkListQueue[T any]() *RLinkListQueue[T] {
	return &RLinkListQueue[T]{
		_list: nil,
	}
}

func (q *RLinkListQueue[T]) Enqueue(v T) {
	temp := q._list
	q._list = rlinkedlist.NewListNode(v)
	q._list.Next = temp
}

func (q *RLinkListQueue[T]) Dequeue(v T) {

}
