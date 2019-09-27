package matrix_ops

import (
	"fmt"
	"strings"
)

type matrix [][]int

func (m matrix) equal(m2 matrix) bool {
	if m2 == nil {
		return false
	}

	if len(m) != len(m2) || len(m[0]) != len(m2[0]) {
		return false
	}

	for i, row := range m {
		for j, val := range row {
			if val != m2[i][j] {
				return false
			}
		}
	}

	return true
}

func (m matrix) square() bool {
	if len(m) != 0 && len(m) == len(m[0]) {
		return true
	}
	return false
}

func (m matrix) String() string {
	b := strings.Builder{}
	for _, row := range m {
		for _, val := range row {
			b.WriteString(fmt.Sprintf("%4d", val))
		}
		b.WriteString("\n")
	}
	return b.String()
}
