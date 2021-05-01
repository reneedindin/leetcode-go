package array

// Version1
func removeDuplicatesV1(nums []int) int {
	idx := 0
	for idx < len(nums)-1 {
		current := nums[idx]
		if current == nums[idx+1] {
			nums = append(nums[:idx+1], nums[idx+2:]...)
			idx = idx
			continue
		}
		idx += 1
	}
	return len(nums)
}

// Version2
func removeDuplicatesV2(nums []int) int {
	count := 0
	only := 0
	for idx := 0; idx < len(nums); idx++ {
		if idx == 0 {
			count ++
		}
		if nums[idx] != nums[only] {
			only ++
			nums[only] = nums[idx]
			count ++
		}
	}

	return count
}
