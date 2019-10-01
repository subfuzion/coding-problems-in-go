package is_unique

import (
	"testing"

	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(string) bool

// Test defines a test case
type Test struct {
	Input    string
	Expected bool
}

// Tests is a test suite
var Tests = []Test{
	{"a", true},
	{"ab", true},
	{"abcdefghijk", true},
	{"abcxyz", true},
	{"aa", false},
	{"abcc", false},
	{"abccde", false},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				if actual := f(test.Input); actual != test.Expected {
					t.Errorf("Expected unique=%t for %s", test.Expected, test.Input)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps descriptive names to solution implementations
var Solutions = []F{
	isUniqueUsingSet,
	isUniqueUsingBitset,
	isUnique1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
Is Unique: Implement an algorithm to determine if a string has all unique
characters.
What if you cannot use additional data structures?
Hints (Cracking the Coding Interview 6): #44, #117, #132
******************************************************************************/
func isUnique1(s string) bool {
	return false
}
