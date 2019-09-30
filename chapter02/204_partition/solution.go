package partition

import ds "github.com/subfuzion/coding-problems-in-go/datastructures"

func partition(n *ds.Node, pivot int) *ds.Node {
	if n == nil {
		return nil
	}

	// insert values < pivot at head
	head := n
	// insert values >= pivot at tail
	tail := n

	for node := n; node != nil; {
		next := node.Next
		if node.Data < pivot {
			node.Next = head
			head = node
		} else {
			node.Next = nil
			tail.Next = node
			tail = node
		}
		node = next
	}

	return head
}
