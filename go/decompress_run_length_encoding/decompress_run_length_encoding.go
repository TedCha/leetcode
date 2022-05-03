package decompress_run_length_encoding

func decompressRLElist(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i += 2 {
		freq := nums[i]
		val := nums[i+1]

		for i := 0; i < freq; i++ {
			result = append(result, val)
		}
	}

	return result
}
