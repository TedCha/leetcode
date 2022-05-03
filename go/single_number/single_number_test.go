package single_number

import "testing"

func Test_Single_Number(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	r := singleNumber(nums)
	if r != 4 {
		t.FailNow()
	}
}
