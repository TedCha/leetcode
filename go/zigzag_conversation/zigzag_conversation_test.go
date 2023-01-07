package zigzag_conversation

import "testing"

func Test_Zigzag_Conversation_1(t *testing.T) {
	result := convert("PAYPALISHIRING", 3)

	if result != "PAHNAPLSIIGYIR" {
		t.FailNow()
	}
}

func Test_Zigzag_Conversation_2(t *testing.T) {
	result := convert("PAYPALISHIRING", 4)

	if result != "PINALSIGYAHRPI" {
		t.FailNow()
	}
}

func Test_Zigzag_Conversation_3(t *testing.T) {
	result := convert("AB", 1)

	if result != "AB" {
		t.FailNow()
	}
}
