package check_permutation

import (
	"sort"
	"strings"
)

func checkPermutationBySort(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	// assume ascii
	achars := strings.Split(a, "")
	bchars := strings.Split(b, "")

	sort.Strings(achars)
	sort.Strings(bchars)

	sa := strings.Join(achars, "")
	sb := strings.Join(bchars, "")

	return sa == sb
}

func checkPermutationByWeight(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	weights := map[rune]int{}

	// for each rune in a, increment its count
	for _, r := range a {
		weights[r]++
	}

	// for each rune in b, decrement its count
	// if we go negative, then there was no corresponding rune
	// in a and we can short circuit out with false
	for _, r := range b {
		weights[r]--
		if weights[r] < 0 {
			return false
		}
	}

	// since the strings were the same length and all letters were accounted
	// for, we can return with success
	return true
}
