package rarrayqueue

type RArrayQueue[T any] struct {
	_array []T
}

func NewRArrayQueue[T any]() *RArrayQueue[T] {
	return &RArrayQueue[T]{
		_array: *new([]T),
	}
}

func (q *RArrayQueue[T]) Enqueue(v T) {
	q._array = append(q._array, v)
}

func (q *RArrayQueue[T]) Dequeue() (v T) {
	v = q._array[0]
	q._array = q._array[1:]
	return
}

func (q *RArrayQueue[T]) Peek() (v T) {
	v = q._array[0]
	return
}

func (q RArrayQueue[T]) Size() int { return len(q._array) }

func (q RArrayQueue[T]) IsEmpty() bool { return q.Size() == 0 }

func (q RArrayQueue[T]) ToSlice() []T { return q._array }
