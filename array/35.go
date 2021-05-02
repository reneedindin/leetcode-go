package array

// version1
func searchInsertV1(nums []int, target int) int {
	for idx := 0; idx < len(nums); idx++ {
		if idx == len(nums)-1 && nums[idx] < target {
			return idx + 1
		}
		if nums[idx] < target {
			continue
		}
		if nums[idx] >= target {
			return idx
		}
	}
	return 0
}

// version2
func searchInsertV2(nums []int, target int) int {
	idx := 0
	for idx < len(nums) {
		if nums[idx] < target {
			idx++
		} else if nums[idx] >= target {
			return idx
		}
	}
	return idx
}