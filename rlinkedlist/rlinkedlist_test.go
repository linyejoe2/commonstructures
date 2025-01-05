package rlinkedlist

import "testing"

func TestInit(t *testing.T) {
	node := NewListNode(0)
	println(node.Val)

	node.Next = NewListNode(9)
	node.Next.Next = NewListNode(20)
	node.Next.Next.Next = NewListNode(29)

	InsertNode(node, NewListNode(23))

	PrintAllVal(node)

	Delete(node, 4)

	PrintAllVal(node)

	println("Access 2: ", Access(node, 2).Val)
}

func TestAccess(t *testing.T) {
	println("test access list with a large int: ", Access(NewListNode(2), 99))
}
