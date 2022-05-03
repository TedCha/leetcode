package string_compression

import (
	"strconv"
)

func compress(chars []byte) int {

	// first byte is first key
	key := chars[0]
	count := 0
	temp := make([]byte, 0)
	for _, c := range chars {
		// if the current byte does not match the key, change the key
		if c != key {
			temp = append(temp, key)
			// only append the count if it's greater than one
			if count > 1 {
				temp = append(temp, []byte(strconv.Itoa(count))...)
			}
			key = c
			count = 0
		}
		count++
	}

	// Get last key/count
	temp = append(temp, key)
	if count > 1 {
		temp = append(temp, []byte(strconv.Itoa(count))...)
	}

	// copy returns the number of elements copied, and copies in place the temp array
	return copy(chars[:len(temp)], temp)
}
