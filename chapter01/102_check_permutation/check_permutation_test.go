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
	Test{"abc", "abc", true},
	Test{"abc", "bca", true},
	Test{"abc", "cab", true},
	Test{"abc", "ab", false},
	Test{"abc", "abcd", false},
	Test{"a", "a", true},
	Test{"a", "b", false},
	Test{"a", "", false},
	Test{"", "b", false},
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
func checkPermutation1(a, b string) bool {
	return false
}
