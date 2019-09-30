package sum_lists

import (
	"github.com/subfuzion/coding-problems-in-go/test"
	"testing"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
)

// F is the type of function that implements the solution
type F func(a, b *ds.Node) *ds.Node

// Test defines a test case
type Test struct {
	A             *ds.Node
	B             *ds.Node
	ExpectReverse *ds.Node
	ExpectForward *ds.Node
}

// Tests is a test suite
var Tests = []Test{
	{
		A:             &ds.Node{7, &ds.Node{1, &ds.Node{6, nil}}},
		B:             &ds.Node{5, &ds.Node{9, &ds.Node{2, nil}}},
		ExpectReverse: &ds.Node{2, &ds.Node{1, &ds.Node{9, nil}}},
		ExpectForward: &ds.Node{1, &ds.Node{3, &ds.Node{0, &ds.Node{8, nil}}}},
	},
	{
		A:             &ds.Node{1, &ds.Node{2, nil}},
		B:             &ds.Node{3, nil},
		ExpectReverse: &ds.Node{4, &ds.Node{2, nil}},
		ExpectForward: &ds.Node{1, &ds.Node{5, nil}},
	},
	{
		A:             &ds.Node{9, nil},
		B:             &ds.Node{9, nil},
		ExpectReverse: &ds.Node{8, &ds.Node{1, nil}},
		ExpectForward: &ds.Node{1, &ds.Node{8, nil}},
	},
}

func TestSolution(t *testing.T) {
	for _, f := range ForwardSolutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := f(test.A, test.B)
				expected := test.ExpectForward
				if !actual.Equal(expected) {
					t.Errorf("\nA: %v\nB: %v\nexpected: %v\nactual:   %v", test.A, test.B, expected, actual)
				}
			}
		})
	}
	for _, f := range ReverseSolutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := f(test.A, test.B)
				expected := test.ExpectReverse
				if !actual.Equal(expected) {
					t.Errorf("\nA: %v\nB: %v\nexpected: %v\nactual:   %v", test.A, test.B, expected, actual)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var ForwardSolutions = []F{
	sumListsForward,
	// sumLists1,
	// ...
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var ReverseSolutions = []F{
	sumListsReverse,
	// sumLists1,
	// ...
}

// TODO - implement your own solution
// Update either the ForwardSolutions or ReverseSolutions list, above,
// as appropriate for your solution.
/*
Sum Lists: You have two numbers represented by a linked list, where each node contains a single digit.
The digits are stored in reverse order, such that the 1 's digit is at the head of the list.
Write a function that adds the two numbers and returns the sum as a linked list.
EXAMPLE
Input:  (7-> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
Output: 2 -> 1 -> 9. That is, 912.
FOLLOW UP
Suppose the digits are stored in forward order. Repeat the above problem.
EXAMPLE
Input: (6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295.
Output: 9 -> 1 -> 2. That is, 912.
 */
func sumLists1(a, b *ds.Node) *ds.Node {
	return nil
}
