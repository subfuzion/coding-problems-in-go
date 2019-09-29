package delete_middle_node

import (
	"testing"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
	"github.com/subfuzion/coding-problems-in-go/test"
)

type F func(node *ds.Node) error

type Test struct {
	// ShouldSucceed is true if a node should be removed
	// ShouldSucceed is false when a node is the end of a list
	// (i.e., when node.Next is nil) and therefore can't be removed
	// (because there is no next node to replace the contents of the
	// supplied node
	ShouldSucceed  bool

	// Input points to a node that is supposed to be in the middle of a list
	Input *ds.Node

	// Expected is what the list should look like after removing a node
	Expected *ds.Node
}

var Tests = []Test{
		{
			ShouldSucceed:  false,
			Input:    &ds.Node{3, nil},
			Expected: &ds.Node{3, nil},
		},
		{
			ShouldSucceed:  true,
			Input: &ds.Node{3, &ds.Node{4, nil}},
			Expected: &ds.Node{4, nil},
		},
		{
			ShouldSucceed:  true,
			Input:   &ds.Node{3, &ds.Node{4, &ds.Node{5, nil}}},
			Expected: &ds.Node{4, &ds.Node{5, nil}},
		},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				// clone the input so that we can still reference the original list for error messages
				list := test.Input.Clone()
				err := f(list)
				expected := test.Expected

				if test.ShouldSucceed && err != nil {
					t.Errorf("error (%s) for input: %v", err, test.Input)
				} else if !test.ShouldSucceed && err == nil {
					t.Errorf("expected nil result for input: %v, but got: %v", test.Input, list)
				}

				if !ds.ListEqual(list, expected) {
					t.Errorf("actual != expected\ninput: %v\nactual: %v\nexpected: %v", test.Input, list, expected)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var Solutions = []F{
	removeMiddle,
	// removeDups1,
	// ...
}

// TODO - implement (and update the Solutions list, above)
/*
Delete Middle Node: Implement an algorithm to delete a node in the middle
(i.e., any node but the first and last node, not necessarily the exact middle)
of a singly linked list, given only access to that node.
EXAMPLE
Input: the node c from the linked list a->b->c->d->e->f
Result: nothing is returned, but the new linked list looks like a->b->d->e->f
 */
func removeMiddle1(node *ds.Node) error {
	return nil
}
