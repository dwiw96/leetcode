package easy

func RemoveDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func RemoveDuplicates2(nums []int) int {
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
