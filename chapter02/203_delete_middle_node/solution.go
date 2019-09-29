package delete_middle_node

import (
	"errors"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
)

func removeMiddle(node *ds.Node) error {
	if node == nil {
		return errors.New("node cannot be nil")
	}
	if node.Next == nil {
		return errors.New("node cannot be the last node")
	}

	node.Data = node.Next.Data
	node.Next = node.Next.Next

	return nil
}