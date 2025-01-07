package rdoublylinkedlist

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
	Prev *ListNode[T]
}

func NewListNode[T any](val T) *ListNode[T] {
	return &ListNode[T]{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}

func InsertNode[T any](head *ListNode[T], insert *ListNode[T]) {
	temp := head.Next
	head.Next = insert
	insert.Next = temp
}

func PrintAllVal[T any](node *ListNode[T]) {
	if node == nil {
		println("")
		return
	}

	print(node.Val, " ")
	PrintAllVal(node.Next)
}

// Can't delete a node when whole list length is 1.
func Delete[T any](node *ListNode[T], depth int) {
	if depth == 0 {
		if node.Next == nil {
			return
		}
		*node = *node.Next
		return
	}

	if depth != 1 {
		Delete(node.Next, depth-1)
		return
	}

	if node.Next == nil {
		return
	}
	node.Next = node.Next.Next
}

func Access[T any](head *ListNode[T], depth int) *ListNode[T] {
	if depth < 0 || head == nil {
		return nil
	}

	if depth == 0 {
		return head
	}

	return Access(head.Next, depth-1)
}
