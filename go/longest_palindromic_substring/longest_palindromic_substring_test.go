package longest_palindromic_substring

import "testing"

func Test_Longest_Palindromic_Substring_Case_1(t *testing.T) {
	s := "babad"
	r := longestPalindrome(s)
	if r != "bab" && r != "aba" {
		t.Fatalf("%s != bab||aba\n", r)
	}
}

func Test_Longest_Palindromic_Substring_Case_2(t *testing.T) {
	s := "cbbd"
	r := longestPalindrome(s)
	if r != "bb" {
		t.Fatalf("%s != bb\n", r)
	}
}

func Test_Longest_Palindromic_Substring_Case_3(t *testing.T) {
	s := "bb"
	r := longestPalindrome(s)
	if r != "bb" {
		t.Fatalf("%s != bb\n", r)
	}
}

func Test_Longest_Palindromic_Substring_Case_4(t *testing.T) {
	s := "aaaa"
	r := longestPalindrome(s)
	if r != "aaaa" {
		t.Fatalf("%s != aaaa\n", r)
	}
}

func Test_Longest_Palindromic_Substring_Case_5(t *testing.T) {
	s := "aacabdkacaa"
	r := longestPalindrome(s)
	if r != "aca" {
		t.Fatalf("%s != aca\n", r)
	}
}

func Test_Longest_Palindromic_Substring_Case_6(t *testing.T) {
	s := "abb"
	r := longestPalindrome(s)
	if r != "bb" {
		t.Fatalf("%s != bb\n", r)
	}
}

func Test_Longest_Palindromic_Substring_Case_7(t *testing.T) {
	s := "ccc"
	r := longestPalindrome(s)
	if r != "ccc" {
		t.Fatalf("%s != bb\n", r)
	}
}
