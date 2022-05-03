package single_number

func singleNumber(nums []int) int {
	r := 0
	for _, num := range nums {
		// evaluate XOR on each int, duplicate numbers will
		// cancel each other out
		r = r ^ num
	}
	return r
}
