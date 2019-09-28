package remove_dups

import ds "github.com/subfuzion/coding-problems-in-go/datastructures"

func removeDupsSet(n *ds.Node) bool {
	set := ds.NewIntSet()

	node := n
	set.Add(node.Data)
	next := node.Next
	for next != nil {
		if set.Add(next.Data) {
			node = next
		}
		next = next.Next
	}
	node.Next = nil
	return true
}

func removeDupsRunner(n *ds.Node) bool {
	for node := n; node != nil; {
		runner := node.Next
		for runner != nil {
			if runner.Data != node.Data {
				break
			} else {
				runner = runner.Next
			}
		}
		node.Next = runner
		node = node.Next
	}

	return true
}
