package rlinkliststack_test

import (
	"testing"

	"github.com/linyejoe2/commonstructures/rstack/rlinkliststack"
	"github.com/stretchr/testify/assert"
)

func TestNewRLinkListStack(t *testing.T) {
	stack := rlinkliststack.NewRLinkListStack[int]()
	assert.NotNil(t, stack)
	assert.True(t, stack.IsEmpty())
	assert.Equal(t, 0, stack.Size())
}

func TestPush(t *testing.T) {
	stack := rlinkliststack.NewRLinkListStack[int]()
	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, 2, stack.Size())
	assert.Equal(t, 2, stack.Peek())
}

func TestPop(t *testing.T) {
	stack := rlinkliststack.NewRLinkListStack[int]()
	stack.Push(1)
	stack.Push(2)

	val := stack.Pop()
	assert.Equal(t, 2, val)
	assert.Equal(t, 1, stack.Size())

	val = stack.Pop()
	assert.Equal(t, 1, val)
	assert.True(t, stack.IsEmpty())

	// Pop from an empty stack should return the zero value
	val = stack.Pop()
	assert.Equal(t, 0, val) // Assuming the zero value for an int is 0
}

func TestPeek(t *testing.T) {
	stack := rlinkliststack.NewRLinkListStack[int]()
	stack.Push(1)
	stack.Push(2)

	val := stack.Peek()
	assert.Equal(t, 2, val)

	stack.Pop()
	val = stack.Peek()
	assert.Equal(t, 1, val)

	stack.Pop()
	// Peek on an empty stack should return the zero value
	val = stack.Peek()
	assert.Equal(t, 0, val)
}

func TestIsEmpty(t *testing.T) {
	stack := rlinkliststack.NewRLinkListStack[int]()
	assert.True(t, stack.IsEmpty())

	stack.Push(1)
	assert.False(t, stack.IsEmpty())

	stack.Pop()
	assert.True(t, stack.IsEmpty())
}

func TestToSlice(t *testing.T) {
	stack := rlinkliststack.NewRLinkListStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	slice := stack.ToSlice()
	assert.Equal(t, []int{3, 2, 1}, slice)

	// Ensure that the stack structure is not affected after ToSlice
	assert.Equal(t, 3, stack.Peek())
}

func TestToSliceWithNewInstance(t *testing.T) {
	stack := rlinkliststack.NewRLinkListStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	slice := stack.ToSliceWithNewInstance()
	assert.Equal(t, &[]int{3, 2, 1}, slice)

	// Ensure that the stack structure is not affected after ToSliceWithNewInstance
	assert.Equal(t, 3, stack.Peek())
}

func TestPopAndPrintAll(t *testing.T) {
	stack := rlinkliststack.NewRLinkListStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Manually checking the printed output might be necessary,
	// or you can use a mock to capture the output.
	stack.PopAndPrintAll()
	assert.True(t, stack.IsEmpty())
}
