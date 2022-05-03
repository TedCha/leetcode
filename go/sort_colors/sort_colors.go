package sort_colors

// quick sort implementation
func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	left, right := 0, len(nums)-1

	// use the middle number in the array as the pivot
	pivot := len(nums) / 2

	// swap the pivot with the right most value
	nums[pivot], nums[right] = nums[right], nums[pivot]

	// places all elements smaller than the pivot to the left of the pivot
	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	// swap the pivot with the last left element
	// to place all numbers greater than the pivot to the right of the pivot
	nums[left], nums[right] = nums[right], nums[left]

	// repeat the process recursively for both the left and right side of the pivot
	sortColors(nums[:left])
	sortColors(nums[left+1:])
}
