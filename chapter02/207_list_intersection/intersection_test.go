//nolint:govet //linter: golangci-lint (disabling warnings for unkeyed fields in composites)
package list_intersection

import (
	"testing"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(list1, list2 *ds.Node) *ds.Node

// Test defines a test case
type Test struct {
	List1    *ds.Node
	List2    *ds.Node
	Expected *ds.Node
}

// Tests is a test suite
var Tests = []Test{}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := f(test.List1, test.List2)
				expected := test.Expected

				if actual != expected {
					t.Errorf("\nexpected %v for intersection of %v and %v", expected, test.List1, test.List2)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var Solutions = []F{
	intersects,
	// intersects1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
Intersection: Given two (singly) linked lists, determine if the two lists
intersect. Return the inter­ secting node. Note that the intersection is
defined based on reference, not value. That is, if the kth node of the first
linked list is the exact same node (by reference) as the jth node of the second
linked list, then they are intersecting.
Hints (Cracking the Coding Interview 6): #20, #45, #55, #65, #76, #93, #111,
                                         #120, #129
******************************************************************************/
func intersects1(list1, list2 *ds.Node) *ds.Node {
	return nil
}
