package zigzag_conversation

import (
	"strings"
)

// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	// create matrix the size of rows to track letters
	matrix := make([][]rune, numRows)

	var result strings.Builder
	count := 1
	additive := 1
	for _, c := range s {

		// oscillate count between min and max (starting row to end row)
		// by using a negative/positive additive
		if count == 1 {
			additive = 1
		} else if count == numRows {
			additive = -1
		}

		// add char to row in which it should be read in zigzag method
		matrix[count-1] = append(matrix[count-1], c)

		count += additive
	}

	// join all of our slices of runes as read order is now set correctly
	for i := range matrix {
		result.WriteString(string(matrix[i]))
	}

	return result.String()
}
