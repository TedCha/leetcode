package longest_palindromic_substring

func longestPalindrome(s string) string {

	var longest string
	for i := range s {
		// perform expand twice on all chars to capture all possible palindromes
		// handles palindromes that are even and odd in length
		// 1st time: expand from one character
		// 2nd time: expand from two characters
		longest = expand(s, i, i, longest)
		longest = expand(s, i, i+1, longest)
	}

	return longest
}

// s - string to traverse
// l - low index
// h - high index
// longest - longest palindrome so far
// expand to the left and right to find all possible palindromes
func expand(s string, l int, h int, longest string) string {
	var p string

	// while between low/high bound and the low bound = high bound
	for l >= 0 && h < len(s) && s[l] == s[h] {
		p = s[l : h+1] // high bound of slice is exclusive
		l--
		h++
	}

	// if the current palindrome is longer than the longest
	if len(longest) < len(p) {
		return p
	}

	return longest
}
