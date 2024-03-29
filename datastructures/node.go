package datastructures

import (
	"strconv"
	"strings"
)

// Node for single linked list of integers
type Node struct {
	Data int
	Next *Node
}

// String stringifies the lists starting from node
func (n Node) String() string {
	b := strings.Builder{}
	b.WriteRune('{')
	next := &n
	for next != nil {
		b.WriteString(strconv.Itoa(next.Data))
		if next.Next != nil {
			b.WriteString(", ")
		}
		next = next.Next
	}
	b.WriteRune('}')
	return b.String()
}

// Equal returns true if n.Data == n2.Data
func (n *Node) Equal(n2 *Node) bool {
	switch {
	case n == n2:
		return true
	case n != nil && n2 != nil:
		return n.Data == n2.Data
	}
	return false
}

// ListEqual compares two lists comprised of a chain of Nodes to check if
// corresponding Data fields are equal for each link in the chain until
// the end of the list is reached.
// List iteration begins with the supplied node pointers and continues
// until reaching the end of either list.
// Returns false if any two nodes aren't equal or if one list is shorter than the other.
func ListEqual(n1 *Node, n2 *Node) bool {
	if n1 == nil || n2 == nil {
		return false
	}

	p1 := n1
	p2 := n2
	for p1.Data == p2.Data {
		p1 = p1.Next
		p2 = p2.Next

		// if either list ends before the other, then lists aren't equal
		if (p1 == nil && p2 != nil) || (p1 != nil && p2 == nil) {
			return false
		}

		// nothing left, so lists are equal
		if p1 == nil && p2 == nil {
			return true
		}
	}

	return false
}

// Length returns length of single linked list starting from n
func (n *Node) Length() int {
	count := 0
	for node := n; node != nil; {
		if node != nil {
			count++
		}
		node = node.Next
	}
	return count
}

// Clone returns a copy of the list with new Node instances
func (n Node) Clone() *Node {
	head := &Node{Data: n.Data}

	for clone, node := head, n.Next; node != nil; {
		clone.Next = &Node{Data: node.Data, Next: node.Next}
		clone = clone.Next
		node = node.Next
	}

	return head
}

// Slice returns a list of nodes as a slice
func (n *Node) Slice() []*Node {
	slice := []*Node{}
	for n != nil {
		slice = append(slice, n)
		n = n.Next
	}
	return slice
}

