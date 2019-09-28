package palindrome

import (
	"testing"

	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(string) bool

// Test defines a test case
type Test struct {
	Input      string
	Palindrome string
	Expected   bool
}

// Tests is a test suite
var Tests = []Test{
	{"atco cta", "taco cat", true},
	{"aan", "ana", true},
	{"anan", "anna", true},
	{"aaanx", "", false},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := isPalindromePermutation(test.Input)
				expected := test.Expected
				if actual != expected {
					if expected {
						t.Errorf("expected '%s' => '%s'", test.Input, test.Palindrome)
					} else {
						t.Errorf("unexpected: '%s'", test.Input)
					}
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps descriptive names to solution implementations
var Solutions = []F{
	isPalindromePermutation,
	// isPalindromePermutation1,
	// ...
}

// TODO - implement (and update the Solutions list, above)
// assumes rune array is large enough to hold the extra characters after conversion
func isPalindromePermutation1(s string) bool {
	return false
}
