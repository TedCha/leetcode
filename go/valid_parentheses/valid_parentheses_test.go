package valid_parentheses

import "testing"

func Test_Valid_Parentheses_Case_1(t *testing.T) {
	input := "()"
	output := isValid(input)
	if !output {
		t.Fatalf("expected true, got %v", output)
	}
}

func Test_Valid_Parentheses_Case_2(t *testing.T) {
	input := "()[]{}"
	output := isValid(input)
	if !output {
		t.Fatalf("expected true, got %v", output)
	}
}

func Test_Valid_Parentheses_Case_3(t *testing.T) {
	input := "(]"
	output := isValid(input)
	if output {
		t.Fatalf("expected false, got %v", output)
	}
}

func Test_Valid_Parentheses_Case_4(t *testing.T) {
	input := "([{}])"
	output := isValid(input)
	if !output {
		t.Fatalf("expected true, got %v", output)
	}
}

func Test_Valid_Parentheses_Case_5(t *testing.T) {
	input := "["
	output := isValid(input)
	if output {
		t.Fatalf("expected false, got %v", output)
	}
}

func Test_Valid_Parentheses_Case_6(t *testing.T) {
	input := "]"
	output := isValid(input)
	if output {
		t.Fatalf("expected false, got %v", output)
	}
}
