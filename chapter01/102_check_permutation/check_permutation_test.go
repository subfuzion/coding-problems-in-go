package check_permutation

import (
	"testing"

	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(string, string) bool

// Test defines a test case
type Test struct {
	A        string
	B        string
	Expected bool
}

// Tests is a test suite
var Tests = []Test{
	{"abc", "abc", true},
	{"abc", "bca", true},
	{"abc", "cab", true},
	{"abc", "ab", false},
	{"abc", "abcd", false},
	{"a", "a", true},
	{"a", "b", false},
	{"a", "", false},
	{"", "b", false},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				if actual := f(test.A, test.B); actual != test.Expected {
					t.Errorf("expected %t for a=%s, b=%s", test.Expected, test.A, test.B)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var Solutions = []F{
	checkPermutationBySort,
	checkPermutationByWeight,
	// checkPermutation1,
	// ...
}

// TODO - implement (and update the Solutions list, above)
/*
Check Permutation: Given two strings,write a method to decide if one is a permutation of the
other.
 */
func checkPermutation1(a, b string) bool {
	return false
}
