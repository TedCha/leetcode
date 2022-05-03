package decompress_run_length_encoding

import "testing"

func Test_Decompress_Run_Length_Encoding_Case_1(t *testing.T) {
	nums := []int{1, 1, 2, 3}
	result := decompressRLElist(nums)
	eq := compareIntSlices(result, []int{1, 3, 3})
	if !eq {
		t.FailNow()
	}
}

func compareIntSlices(a []int, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
