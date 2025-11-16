package main

// removeDuplicates
// Given an integer array nums sorted in non-decreasing order, remove the
// duplicates in-place such that each unique element appears only once. The
// relative order of the elements should be kept the same. Then return the number
// of unique elements in nums.
func removeDuplicates(nums []int) int {
	if l := len(nums); l <= 1 {
		return l
	}

	current := nums[0]
	uniqueCount := 1

	for _, e := range nums[1:] {
		if e != current {
			// e is not a repeated value
			nums[uniqueCount] = e
			uniqueCount += 1
			current = e
		}
	}

	return uniqueCount
}
