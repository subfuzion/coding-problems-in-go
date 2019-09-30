package palindrome

import ds "github.com/subfuzion/coding-problems-in-go/datastructures"

func isPalindrome(n *ds.Node) bool {
	if n == nil {
		return false
	}

	rev, length := reverseClone(n)

	for i := 0; i < length/2; i++ {
		if n.Data != rev.Data {
			return false
		}
		n = n.Next
		rev = rev.Next
	}

	return true
}

func reverseClone(n *ds.Node) (*ds.Node, int) {
	var f func(*ds.Node) (*ds.Node, *ds.Node, int)
	f = func(n *ds.Node) (*ds.Node, *ds.Node, int) {
		if n == nil {
			return nil, nil, 0
		}

		clone := &ds.Node{Data: n.Data}
		head, node, count := f(n.Next)
		if node == nil {
			node = clone
			head = node
		} else {
			node.Next = clone
		}
		return head, clone, count + 1
	}

	// we needed to keep track of the current clone node, during recursion
	// but all we really need to return is the new head and the length
	head, _, length := f(n)
	return head, length
}