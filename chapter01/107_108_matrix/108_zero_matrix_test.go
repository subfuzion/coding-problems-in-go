package matrix_ops

import (
	"testing"

	"github.com/subfuzion/coding-problems-in-go/pkg/test"
)

// ZeroF is the type of function that implements the solution
type ZeroF func(matrix) bool

// ZeroTest defines a test case
type ZeroTest struct {
	Input    matrix
	Expected matrix
}

var ZeroTests = []ZeroTest{
	{
		Input: matrix{
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
		},
		Expected: matrix{
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
		},
	},
	{
		Input: matrix{
			{1, 1, 1, 1, 1, 1},
			{1, 1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 1},
			{1, 1, 1, 1, 1, 1},
		},
		Expected: matrix{
			{1, 1, 0, 1, 0, 1},
			{0, 0, 0, 0, 0, 0},
			{1, 1, 0, 1, 0, 1},
			{0, 0, 0, 0, 0, 0},
			{1, 1, 0, 1, 0, 1},
		},
	},
	{
		Input: matrix{
			{1, 0, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
		},
		Expected: matrix{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0},
		},
	},
	{
		Input: matrix{
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
		},
		Expected: matrix{
			{0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
		},
	},
}

func TestZeroSolution(t *testing.T) {
	for _, f := range ZeroSolutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range ZeroTests {
				actual := test.Input
				expected := test.Expected

				// if valid input then matrix is zerod in place
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
var ZeroSolutions = []ZeroF{
	nullify,
	// nullify1,
	// ...
}

// TODO - implement (and update the Solutions list, above)
// assumes rune array is large enough to hold the extra characters after conversion
func nullify1(m matrix) bool {
	return false
}
