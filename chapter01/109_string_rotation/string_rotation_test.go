package string_rotation

import (
	"testing"

	"github.com/subfuzion/coding-problems-in-go/pkg/test"
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
	{"a", "a", true},
	{"ab", "ab", true},
	{"ab", "ba", true},
	{"abc", "bca", true},
	{"abc", "cab", true},
	{"waterbottle", "erbottlewat", true},
	{"abc", "cba", false},
	{"", "", false},
	{"ab", "abc", false},
	{"abc", "xyz", false},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := isStringRotation(test.A, test.B)
				expected := test.Expected

				if actual != expected {
					t.Errorf("'%s' is a rotation of '%s' should be %t", test.B, test.A, test.Expected)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps a name (or brief description) to solution implementations
var Solutions = []F{
	isStringRotation,
	// isStringRotation1,
	// ...
}

// TODO - implement (and update the Solutions list, above)
func isStringRotation1(a, b string) bool {
	return false
}
