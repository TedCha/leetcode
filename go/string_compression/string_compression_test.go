package string_compression

import "testing"

func Test_String_Compression_Case_1(t *testing.T) {
	chars := []byte("aabbccc")
	r := compress(chars)
	eq := compareByteSlices(chars, []byte("a2b2c3"))
	if r != 6 || !eq {
		t.FailNow()
	}
}

func Test_String_Compression_Case_2(t *testing.T) {
	chars := []byte("a")
	r := compress(chars)
	eq := compareByteSlices(chars, []byte("a"))
	if r != 1 || !eq {
		t.FailNow()
	}
}

func Test_String_Compression_Case_3(t *testing.T) {
	chars := []byte("abbbbbbbbbbbb")
	r := compress(chars)
	eq := compareByteSlices(chars, []byte("ab12"))
	if r != 4 || !eq {
		t.FailNow()
	}
}

func compareByteSlices(a []byte, b []byte) bool {
	// only need to compare the first n expected digits of the array
	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
