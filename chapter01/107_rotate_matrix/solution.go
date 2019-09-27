package matrix_ops

import "github.com/subfuzion/coding-problems-in-go/datastructures"

func rotate(m datastructures.Matrix) bool {
	if !m.Square() {
		return false
	}

	// work our way layer by layer toward the middle
	n := len(m)
	for layer := 0; layer < n/2; layer++ {
		// this is a square matrix, so first and last identify starting and ending
		// boundaries of layer; ex (clockwise from top left):
		// [first][first], [first][last], [last][last], [last][first]
		// Ex: 6x6 matrix
		// iter 0:  first=0 last=5
		// iter 1:  first=1 last=4
		// iter 2:  first=2 last=3
		first := layer
		last := n - 1 - layer

		for i := first; i < last; i++ {
			// As we advance, i represents the current column on the top row and
			// the current row on the right side. Because we are rotating, the
			// bottom and left sides advance in a negative direction, so we need
			// a zero-based offset to compute corresponding locations from the ends.
			// For every rotation, the cells that being shifted are identified as:
			// (given the boundaries: [first][first], [last][last])
			// top side:    [first][i] (advances right)
			// left side:   [last-offset][first] (advances up)
			// bottom side: [last][last-offset] (advances left)
			// right side:  [i][last] (advances down)
			offset := i - first

			// save layer top cell before overwriting with corresponding cell from left side
			// assume saving layer top row cells first, from first to last for each i
			top := m[first][i]

			// move left to top
			m[first][i] = m[last-offset][first]

			// move bottom to left
			m[last-offset][first] = m[last][last-offset]

			// move right to bottom
			m[last][last-offset] = m[i][last]

			// move (saved) top to right
			m[i][last] = top
		}
	}
	return true
}
