package rlist

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	list := NewList[int](10, 2)
	println("Size(): ", list.Size())

	list.Set(20, 4)
	println("New size after Set(20, 4): ", list.Size())
	println("Get(4): ", list.Get(4))

	list.Push(342)
	println("New size after Push(342): ", list.Size())
	println("Get(list.Size()-1): ", list.Get(list.Size()-1))

	list.Insert(2345, list._length-1)
	list.Insert(1, 0)
	fmt.Printf("%v\n", list._arr)

	list.Delete(0)
	fmt.Printf("%v\n", list._arr)
}
