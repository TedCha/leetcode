package palindrome_number

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	var reversed int
	temp := x
	for temp > 0 {
		// temp is parsed by hundreths and reverse builds by hundreths
		reversed = reversed*10 + temp%10
		temp /= 10
	}

	return x == reversed
}
