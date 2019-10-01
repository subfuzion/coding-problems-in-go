//nolint:govet //linter: golangci-lint (disabling warnings for unkeyed fields in composites)
package partition

import (
	"testing"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(n *ds.Node, pivot int) *ds.Node

// Test defines a test case
type Test struct {
	ID       int
	Input    *ds.Node
	Pivot    int
	TestFunc func(*ds.Node, int) bool
}

func testFunc(n *ds.Node, pivot int) bool {
	pivoted := false

	for next := n; next != nil; {
		if next.Data >= pivot {
			pivoted = true
		}
		if next.Data < pivot && pivoted {
			return false
		}
		next = next.Next
	}

	return true
}

// Tests is a test suite
var Tests = []Test{
	{
		ID:       1,
		Input:    &ds.Node{1, &ds.Node{2, &ds.Node{3, nil}}},
		Pivot:    1,
		TestFunc: testFunc,
	},
	{
		ID:       2,
		Input:    &ds.Node{1, &ds.Node{2, &ds.Node{3, nil}}},
		Pivot:    2,
		TestFunc: testFunc,
	},
	{
		ID:       3,
		Input:    &ds.Node{1, &ds.Node{2, &ds.Node{3, nil}}},
		Pivot:    3,
		TestFunc: testFunc,
	},
	{
		ID:       4,
		Input:    &ds.Node{3, &ds.Node{8, &ds.Node{5, &ds.Node{10, &ds.Node{2, &ds.Node{1, nil}}}}}},
		Pivot:    5,
		TestFunc: testFunc,
	},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := f(test.Input.Clone(), test.Pivot)

				// to see log output, use -v
				// $ go test -v -timeout 30s -run TestSolution
				t.Logf("test #%d\nInput    : %v\nPartioned: %v", test.ID, test.Input, actual)

				if ok := test.TestFunc(actual, test.Pivot); !ok {
					t.Errorf("test #%d failed", test.ID)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var Solutions = []F{
	partition,
	// partition1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
Partition: Write code to partition a linked list around a value x, such that
all nodes less than xcome before all nodes greater than or equal to x. If x is
contained within the list, the values of x only need to be after the elements
less than x (see below). The partition element x can appear anywhere in the
"right partition"; it does not need to appear between the left and right
partitions.
EXAMPLE
Input:  3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition=5]
Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
Hints (Cracking the Coding Interview 6): #3, #24
******************************************************************************/
func partition1(n *ds.Node, pivot int) *ds.Node {
	return nil
}
