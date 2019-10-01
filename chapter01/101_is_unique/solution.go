package is_unique

//
// Solution using a set
//

// isUniqueUsingSet uses a set (implemented with a map) to check
// if characters in a string are unique.
func isUniqueUsingSet(s string) bool {
	// assume ASCII, but this can be modified for any character set
	const sizeCharset = 128

	// If s is either empty or if there are more characters than the
	// size of the character set, then somewhere in the string there
	// must be duplicate characters
	if s == "" || len(s) > sizeCharset {
		return false
	}

	// map of runes to empty structs (the empty structs serve as flags)
	// see: https://dave.cheney.net/2014/03/25/the-empty-struct
	chars := map[rune]struct{}{/* empty map to start */}

	// for each rune in s
	for _, c := range s {
		if _, found := chars[c]; found {
			return false
		}
		//
		chars[c] = struct{}{/* empty struct takes no space */}
	}

	return true
}

//
// Solution using a bitset
//

// isUniqueUsingBitset uses a bitset to check if characters in a
// string are unique. There must be at least as many bits as there
// are characters in the character set (this solution assumes only ASCII).
// This solution dispenses with the check shown in the previous one
// for whether s either is empty or has more characters than the character
// set.
func isUniqueUsingBitset(s string) bool {
	chars := 0

	for _, c := range s {
		val := c - 'a'
		mask := 1 << val
		if chars & mask > 0 {
			return false
		}
		chars |= mask
	}

	return true
}
