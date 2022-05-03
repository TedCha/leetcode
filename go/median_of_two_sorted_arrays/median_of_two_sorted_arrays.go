package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// create merged list sorted result slice
	result := make([]int, len(nums1)+len(nums2))

	// create our result until we have reached the maximum elems of one of the slices
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			result[k] = nums1[i]
			k++
			i++
		} else {
			result[k] = nums2[j]
			k++
			j++
		}
	}

	// x - the leftover nums in the nums1 slice
	// y - the leftover nums in the nums2 slice
	// k - how much of our result slice we covered,
	x := k + len(nums1) - i // i - how much of our nums1 slice we covered
	y := k + len(nums2) - j // y - how much of our nums2 slice we covered
	copy(result[k:x], nums1[i:])
	copy(result[k:y], nums2[j:])

	var m float64
	if len(result)%2 == 0 {
		m = (float64(result[len(result)/2]) + float64(result[(len(result)/2)-1])) / 2
	} else {
		m = float64(result[len(result)/2])
	}

	return m
}
