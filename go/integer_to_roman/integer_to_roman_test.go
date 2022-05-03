package integer_to_roman

import "testing"

func Test_Integer_To_Roman_Case_1(t *testing.T) {
	input := 3
	expec := "III"
	result := intToRoman(input)
	if result != expec {
		t.Fatalf("%s does not equal %s\n", result, expec)
	}
}

func Test_Integer_To_Roman_Case_2(t *testing.T) {
	input := 58
	expec := "LVIII"
	result := intToRoman(input)
	if result != expec {
		t.Fatalf("%s does not equal %s\n", result, expec)
	}
}

func Test_Integer_To_Roman_Case_3(t *testing.T) {
	input := 1994
	expec := "MCMXCIV"
	result := intToRoman(input)
	if result != expec {
		t.Fatalf("%s does not equal %s\n", result, expec)
	}
}
