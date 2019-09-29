package get_kth_to_last

import ds "github.com/subfuzion/coding-problems-in-go/datastructures"

// kthToLastRecursive returns kth node
// for k=0, kth node is nil; for k=1, kth node is last item in list;
// k=2, second to last item; etc.
func kthToLastRecursive(n *ds.Node, k int) *ds.Node {
	length := 0

	var f func(n *ds.Node, k int) (*ds.Node, int)
	f = func(n *ds.Node, k int) (*ds.Node, int) {
		// base case
		if n == nil {
			return nil, 0
		}

		// keep returning current node until length reaches k
		kth, length := f(n.Next, k)
		if length < k {
			return n, length + 1
		}

		// don't update kth anymore now that the kth node has been identified
		// no real need to continue incrementing length, but might as well return
		// something useful (ie, full length of the list traversed)
		return kth, length + 1
	}

	// make sure list was actually long enough for kth node
	kth, length := f(n, k)
	if length < k {
		return nil
	}
	return kth
}

// Non-recursive version
func kthToLastIterative(n *ds.Node, k int) *ds.Node {
	if n == nil || k == 0 {
		return nil
	}

	p2 := n
	for i := 0; i < k; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}

	for p2 != nil {
		n = n.Next
		p2 = p2.Next
	}

	return n
}