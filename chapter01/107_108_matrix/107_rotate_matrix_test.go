package matrix_ops

import (
	"testing"

	"github.com/subfuzion/coding-problems-in-go/pkg/test"
)

// RotateF is the type of function that implements the solution
type RotateF func(matrix) bool

// RotateTest defines a test case
type RotateTest struct {
	Input    matrix
	Expected matrix
}

var RotateTests = []RotateTest{
	{
		Input: matrix{
			{0, 1, 2},
			{3, 4, 5},
			{6, 7, 8},
		},
		Expected: matrix{
			{6, 3, 0},
			{7, 4, 1},
			{8, 5, 2},
		},
	},
	{
		Input: matrix{
			{0, 1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10, 11},
			{12, 13, 14, 15, 16, 17},
			{18, 19, 20, 21, 22, 23},
			{24, 25, 26, 27, 28, 29},
			{30, 31, 32, 33, 34, 35},
		},
		Expected: matrix{
			{30, 24, 18, 12, 6, 0},
			{31, 25, 19, 13, 7, 1},
			{32, 26, 20, 14, 8, 2},
			{33, 27, 21, 15, 9, 3},
			{34, 28, 22, 16, 10, 4},
			{35, 29, 23, 17, 11, 5},
		},
	},
}

func TestRotateSolution(t *testing.T) {
	for _, f := range RotateSolutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range RotateTests {
				actual := test.Input
				expected := test.Expected

				// if valid input then matrix is rotated in place
				ok := f(test.Input)

				if !ok {
					t.Errorf("error rotating matrix (must be square with non-zero dimensions):\n%v", test.Input)
				}

				if !actual.equal(expected) {
					t.Errorf("actual != expected\nActual:\n%sExpected:\n%s", actual, expected)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps descriptive names to solution implementations
var RotateSolutions = []RotateF{
	rotate,
	// rotate1,
	// ...
}

// TODO - implement (and update the Solutions list, above)
// assumes rune array is large enough to hold the extra characters after conversion
func rotate1(m matrix) bool {
	return false
}
