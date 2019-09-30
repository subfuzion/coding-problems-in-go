//nolint:govet //linter: golangci-lint (disabling warnings for unkeyed fields in composites)
package remove_dups

import (
	"testing"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(*ds.Node) bool

// Test defines a test case
type Test struct {
	Input    *ds.Node
	Expected *ds.Node
}

var listNoDuplicates = &ds.Node{
	Data: 1, Next: &ds.Node{
		Data: 2, Next: &ds.Node{
			Data: 3, Next: nil,
		},
	},
}

// Tests is a test suite
var Tests = []Test{
	{
		// no duplicates: 1->2->3
		Input: &ds.Node{
			Data: 1, Next: &ds.Node{
				Data: 2, Next: &ds.Node{
					Data: 3, Next: nil,
				},
			},
		},
		Expected: listNoDuplicates,
	},
	{
		// duplicates: 1->*1*->2->3
		Input: &ds.Node{
			Data: 1, Next: &ds.Node{
				Data: 1, Next: &ds.Node{
					Data: 2, Next: &ds.Node{
						Data: 3, Next: nil,
					},
				},
			},
		},
		Expected: listNoDuplicates,
	},
	{
		// duplicates: 1->2->*2*->3
		Input: &ds.Node{
			Data: 1, Next: &ds.Node{
				Data: 2, Next: &ds.Node{
					Data: 2, Next: &ds.Node{
						Data: 3, Next: nil,
					},
				},
			},
		},
		Expected: listNoDuplicates,
	},
	{
		// duplicates: 1->2->3->*3*
		Input: &ds.Node{
			Data: 1, Next: &ds.Node{
				Data: 2, Next: &ds.Node{
					Data: 3, Next: &ds.Node{
						Data: 3, Next: nil,
					},
				},
			},
		},
		Expected: listNoDuplicates,
	},
}

func TestSolution(t *testing.T) {

	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := test.Input
				expected := test.Expected

				if ok := f(actual); !ok {
					t.Errorf("problem with input: %v", actual)
				}

				if !ds.ListEqual(actual, expected) {
					t.Errorf("\nactual:   %v\nexpected: %v", actual, expected)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var Solutions = []F{
	removeDupsSet,
	removeDupsRunner,
	// removeDups1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
Remove Dups: Write code to remove duplicates from an unsorted linked list. How
would you solve this problem if a temporary buffer is not allowed?
******************************************************************************/
func removeDupsSet1(n *ds.Node) bool {
	return false
}
