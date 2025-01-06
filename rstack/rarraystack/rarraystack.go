package rarraystack

import (
	"slices"
)

type RArrayStack[T any] struct {
	_array []T
}

func NewRArrayStack[T any]() *RArrayStack[T] {
	return &RArrayStack[T]{
		_array: nil,
	}
}

func (s *RArrayStack[T]) Push(val T) {
	s._array = append(s._array, val)
}

func (s *RArrayStack[T]) Pop() (val T) {
	if s._array == nil || len(s._array) == 0 {
		return
	}

	val = s._array[s.Size()-1]
	s._array = s._array[:s.Size()-1]
	return
}

func (s *RArrayStack[T]) Peek() T {
	if len(s._array) == 0 {
		return *new(T)
	}
	return s._array[s.Size()-1]
}

func (s *RArrayStack[T]) Size() int { return len(s._array) }

func (s *RArrayStack[T]) IsEmpty() bool { return s.Size() == 0 }

func (s *RArrayStack[T]) ToSlice() []T {
	temp := slices.Clone(s._array)
	slices.Reverse(temp)
	return temp
}
