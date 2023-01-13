package palindrome_number

import "testing"

func Test_Palindrom_Number_1(t *testing.T) {
	result := isPalindrome(121)

	if result != true {
		t.FailNow()
	}
}

func Test_Palindrom_Number_2(t *testing.T) {
	result := isPalindrome(1221)

	if result != true {
		t.FailNow()
	}
}

func Test_Palindrom_Number_3(t *testing.T) {
	result := isPalindrome(2221)

	if result != false {
		t.FailNow()
	}
}

func Test_Palindrom_Number_4(t *testing.T) {
	result := isPalindrome(1231)

	if result != false {
		t.FailNow()
	}
}

func Test_Palindrom_Number_5(t *testing.T) {
	result := isPalindrome(1321)

	if result != false {
		t.FailNow()
	}
}
