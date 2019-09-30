package palindrome

import (
	"testing"

	ds "github.com/subfuzion/coding-problems-in-go/datastructures"
	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(n *ds.Node) bool

// Test defines a test case
type Test struct {
	Input    *ds.Node
	Expected bool}

// Tests is a test suite
var Tests = []Test{
	{
		Input:    &ds.Node{1, &ds.Node{2, &ds.Node{1, nil}}},
		Expected: true,
	},
	{
		Input:    &ds.Node{1, &ds.Node{2, &ds.Node{2, &ds.Node{1, nil}}}},
		Expected: true,
	},
	{
		Input:    &ds.Node{1, nil},
		Expected: true,
	},
	{
		Input:    &ds.Node{1, &ds.Node{2, &ds.Node{2, &ds.Node{2, nil}}}},
		Expected: false,
	},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := f(test.Input)
				expected := test.Expected

				if actual != expected {
					t.Errorf("\nexpected %t for %v", expected, test.Input)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var Solutions = []F{
	isPalindrome,
	// isPalindrome1,
	// ...
}

// TODO - implement your own solution (and update the Solutions list, above)
/*
Palindrome: Implement a function to check if a linked list is a palindrome.
 */
func isPalindrome1(n *ds.Node) bool {
	return false
}
