package rarrayqueue_test

import (
	"testing"

	"github.com/linyejoe2/commonstructures/rqueue/rarrayqueue"
	"github.com/stretchr/testify/assert"
)

func TestNewRLinkListQueue(t *testing.T) {
	q := rarrayqueue.NewRArrayQueue[float64]()
	assert.NotNil(t, q)
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.IsEmpty())
}

func TestEnqueue(t *testing.T) {
	q := rarrayqueue.NewRArrayQueue[float64]()
	q.Enqueue(1.1)
	q.Enqueue(2.2)
	q.Enqueue(3.3)

	assert.Equal(t, 3, q.Size())
	assert.Equal(t, 1.1, q.Peek())
	assert.False(t, q.IsEmpty())
}

func TestDequeue(t *testing.T) {
	q := rarrayqueue.NewRArrayQueue[float64]()
	q.Enqueue(1.1)
	q.Enqueue(2.2)

	val := q.Dequeue()
	assert.Equal(t, 1.1, val)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 2.2, q.Peek())

	val = q.Dequeue()
	assert.Equal(t, 2.2, val)
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.IsEmpty())
}

func TestPeek(t *testing.T) {
	q := rarrayqueue.NewRArrayQueue[float64]()
	q.Enqueue(1.1)
	q.Enqueue(2.2)

	val := q.Peek()
	assert.Equal(t, 1.1, val)
	assert.Equal(t, 2, q.Size()) // Ensure Peek doesn't modify the queue
}

func TestIsEmpty(t *testing.T) {
	q := rarrayqueue.NewRArrayQueue[float64]()
	assert.True(t, q.IsEmpty())

	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}

func TestToSlice(t *testing.T) {
	q := rarrayqueue.NewRArrayQueue[float64]()
	q.Enqueue(float64(1))
	q.Enqueue(float64(2))
	q.Enqueue(float64(3))

	slice := q.ToSlice()
	expectedSlice := []float64{1, 2, 3}
	assert.Equal(t, expectedSlice, slice)
}
