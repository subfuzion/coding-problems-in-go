package get_kth_to_last

import (
	"testing"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
	"github.com/subfuzion/coding-problems-in-go/test"
)

type F func(*ds.Node, int) *ds.Node

type Test struct {
	Input    *ds.Node
	K        int
	Expected *ds.Node
}

var list1 = &ds.Node{1, nil}
var list4 = &ds.Node{1, &ds.Node{2, &ds.Node{3, &ds.Node{4, nil}}}}

var Tests = []Test{
	{
		Input:    list1,
		K:        0,
		Expected: nil,
	},
	{
		Input:    list1,
		K:        1,
		Expected: &ds.Node{1, nil},
	},
	{
		Input:    list1,
		K:        2,
		Expected: nil,
	},
	{
		Input:    list4,
		K:        0,
		Expected: nil,
	},
	{
		Input:    list4,
		K:        1,
		Expected: &ds.Node{4, nil},
	},
	{
		Input:    list4,
		K:        2,
		Expected: &ds.Node{3, &ds.Node{4, nil}},
	},
	{
		Input:    list4,
		K:        4,
		Expected: &ds.Node{1, &ds.Node{2, nil}},
	},
	{
		Input:    list4,
		K:        5,
		Expected: nil,
	},
}

func TestSolution(t *testing.T) {

	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := f(test.Input, test.K)
				expected := test.Expected

				if !actual.Equal(expected) {
					t.Errorf("k=%d\ninput:\n%v\nexpected:\n%v\nactual:\n%v", test.K, test.Input, expected, actual)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var Solutions = []F{
	kthToLastRecursive,
	kthToLastIterative,
	// kthToLast1,
	// ...
}

// TODO - implement (and update the Solutions list, above)
func kthToLast1(n *ds.Node, k int) *ds.Node {
	return nil
}
