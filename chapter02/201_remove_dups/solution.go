package remove_dups

import ds "github.com/subfuzion/coding-problems-in-go/datastructures"

func removeDupsSet(n *ds.Node) bool {
	set := ds.NewIntSet()

	node := n
	set.Add(node.Data)
	for node.Next != nil {
		if set.Add(node.Next.Data) {
			node = node.Next
		} else {
			node.Next = node.Next.Next

		}
	}
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
