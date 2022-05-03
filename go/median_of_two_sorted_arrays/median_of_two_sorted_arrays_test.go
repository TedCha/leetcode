package median_of_two_sorted_arrays

import "testing"

func Test_Median_Of_Two_Sorted_Arrays_Case_1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	r := findMedianSortedArrays(nums1, nums2)

	if r != 2.0 {
		t.FailNow()
	}
}

func Test_Median_Of_Two_Sorted_Arrays_Case_2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	r := findMedianSortedArrays(nums1, nums2)

	if r != 2.5 {
		t.FailNow()
	}
}
