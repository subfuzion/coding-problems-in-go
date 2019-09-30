package is_unique

//
// Solution using a set
//

// RuneSet maps runes to empty struct{}'s
// https://dave.cheney.net/2014/03/25/the-empty-struct
type RuneSet map[rune]struct{}

func NewRuneSet() *RuneSet {
	return &RuneSet{}
}

// Add r to the set and return true; if already present in
// the set then return false.
func (set *RuneSet) Add(r rune) bool {
	if _, exists := (*set)[r]; exists {
		return false
	}
	(*set)[r] = struct{}{}
	return true
}

// isUniqueUsingSet uses a set to check if characters in a string are unique.
func isUniqueUsingSet(s string) bool {
	set := NewRuneSet()
	for _, r := range s {
		if ok := set.Add(r); !ok {
			return false
		}
	}
	return true
}

//
// Solution using a bitset
//

// isUniqueUsingBitset uses a bitset to check if characters in a
// string are unique. There must be at least as many bits as there
// are characters in the character set (this solution assumes only ASCII).
func isUniqueUsingBitset(s string) bool {
	set := 0
	for _, ch := range s {
		val := ch - 'a'
		mask := 1 << uint32(val)
		if set&mask > 0 {
			return false
		}
		set |= mask
	}
	return true
}
