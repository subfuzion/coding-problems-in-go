package sum_lists

import 	ds "github.com/subfuzion/coding-problems-in-go/datastructures"

func sumListsReverse(a, b *ds.Node) *ds.Node {
	if a == nil || b == nil {
		return nil
	}

	carry := 0
	var head *ds.Node
	var tail *ds.Node

	pa := a
	pb := b

	for pa != nil && pb != nil {
		sum := pa.Data + pb.Data + carry
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		}
		n := &ds.Node{sum, nil}
		if head == nil {
			head = n
			tail = n
		} else {
			tail.Next = n
			tail = n
		}
		pa = pa.Next
		pb = pb.Next
	}

	var pr *ds.Node
	if pa != nil {
		pr = pa
	} else if pb != nil {
		pr = pb
	}
	for pr != nil {
		sum := pr.Data + carry
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		}
		n := &ds.Node{sum, nil}
		tail.Next = n
		tail = n
		pr = pr.Next
	}

	if carry > 0 {
		n := &ds.Node{carry, nil}
		tail.Next = n
		tail = n
	}

	return head
}

func sumListsForward(a, b *ds.Node) *ds.Node {
	if a == nil || b == nil {
		return nil
	}

	la := a.Length()
	lb := b.Length()
	if la != lb {
		var p **ds.Node
		var nPadding int
		if la < lb {
			p = &a
			nPadding = lb - la
		} else {
			p = &b
			nPadding = la - lb
		}
		for i := 0; i < nPadding; i++ {
			n := &ds.Node{0, *p}
			*p = n
		}
	}

	var f func(a, b *ds.Node) (*ds.Node, int)
	f = func(a, b *ds.Node) (*ds.Node, int) {
		if a == nil {
			return nil, 0
		}

		next, carry := f(a.Next, b.Next)

		sum := a.Data + b.Data + carry
		if sum > 10 {
			carry = sum / 10
			sum = sum % 10
		}
		n := &ds.Node{sum, next}
		return n, carry
	}

	c, carry := f(a, b)
	if carry > 0 {
		n := &ds.Node{carry, c}
		c = n
	}
	return c
}

