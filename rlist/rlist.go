package rlist

type List[T any] struct {
	_capacity    int
	_length      int
	_arr         []T
	_extendRatio int
}

func NewList[T any](capacity int, extendRatio int) *List[T] {
	return &List[T]{
		_capacity:    capacity,
		_length:      0,
		_arr:         make([]T, capacity),
		_extendRatio: extendRatio,
	}
}

func (l *List[T]) Size() int {
	return l._length
}

func (l *List[T]) Capacity() int {
	return l._capacity
}

func (l *List[T]) Set(val T, index int) {
	if index < 0 || index > l._capacity-1 {
		panic("Index out of range!")
	}

	l._arr[index] = val

	l._length = max(l._length, index+1)
}

func (l *List[T]) Get(index int) T {
	if index < 0 || index > l._capacity-1 {
		panic("Index out of range!")
	}

	return l._arr[index]
}

func (l *List[T]) Push(e T) {
	if l._length == l._capacity {
		l.ExtendCapacity()
	}

	l._arr[l._length] = e
	l._length++
}

func (l *List[T]) Insert(e T, index int) {
	if index < 0 || index > l._capacity-1 {
		panic("Index out of range!")
	}

	if l._length == l._capacity {
		l.ExtendCapacity()
	}

	l._arr = append(append(l._arr[:index], e), l._arr[index+1:]...)

	l._length++
}

func (l *List[T]) Delete(index int) {
	if index < 0 || index > l._capacity-1 {
		panic("Index out of range!")
	}

	l._arr = append(l._arr[:index], l._arr[index+1:]...)
}

func (l *List[T]) ExtendCapacity() {
	l._arr = append(l._arr, make([]T, l._capacity*(l._extendRatio-1))...)

	l._capacity = len(l._arr)
}

func (l *List[T]) ToArray() []T {
	return l._arr[:l._length]
}
