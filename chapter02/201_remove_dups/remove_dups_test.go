package remove_dups

import (
	"github.com/subfuzion/coding-problems-in-go/test"
	"testing"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
)

type F func(*ds.Node) bool

type Test struct {
	Input    *ds.Node
	Expected *ds.Node
}

var Tests = []Test{
	{
		&ds.Node{1, &ds.Node{2, &ds.Node{3, nil}}},
		&ds.Node{1, &ds.Node{2, &ds.Node{3, nil}}},
	},
	{
		&ds.Node{1, &ds.Node{2, &ds.Node{3, &ds.Node{3, nil}}}},
		&ds.Node{1, &ds.Node{2, &ds.Node{3, nil}}},
	},
}

func TestSolution(t *testing.T) {

	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				// we want to be able to reuse the sample data, so clone it
				actual := test.Input.Clone()
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

// TODO - implement (and update the Solutions list, above)
func removeDupsSet1(n *ds.Node) bool {
	return false
}
