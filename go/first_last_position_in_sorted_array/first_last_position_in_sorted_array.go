package first_last_position_in_sorted_array

func searchRange(nums []int, target int) []int {
	isTargetSearch := false
	start := -1
	end := -1

	for i, n := range nums {
		if !isTargetSearch && n == target {
			start = i
			isTargetSearch = true
		}

		if isTargetSearch {
			if i == len(nums)-1 && n == target {
				end = i
				isTargetSearch = false
			} else if n != target {
				end = i - 1
				isTargetSearch = false
			}

		}
	}

	return []int{start, end}
}
