//nolint:govet //linter: golangci-lint (disabling warnings for unkeyed fields in composites)
package loop_detection

import (
	"testing"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(*ds.Node) *ds.Node

// Test defines a test case
type Test struct {
	Input    *ds.Node
	Expected *ds.Node
}

// Tests is a test suite
var Tests = []Test{}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := f(test.Input)
				expected := test.Expected

				if actual != expected {
					t.Errorf("\nexpected %v for %v", expected, test.Input)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var Solutions = []F{
	detectLoop,
	// detectLoop1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
Loop Detection: Given a circular linked list, implement an algorithm that
returns the node at the beginning of the loop.
DEFINITION
Circular linked list: A (corrupt) linked list in which a node's next pointer
points to an earlier node, so as to make a loop in the linked list.
EXAMPLE
Input:  A -> B -> C -> D -> E -> C [the same C as earlier]
Output: C
******************************************************************************/
func detectLoop1(node *ds.Node) *ds.Node {
	return nil
}
