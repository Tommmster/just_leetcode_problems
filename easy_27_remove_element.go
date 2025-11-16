package main

// removeElement
// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The order of the elements may be changed.
// Then return the number of elements in nums which are not equal to val.
func removeElement(nums []int, val int) int {

	valCount := 0

	// Change the array nums such that the first k elements of
	// nums contain the elements which are not equal to val. The
	// remaining elements of nums are not important as well as
	// the size of nums.
	l := len(nums)

	for i := l - 1; i >= 0; i-- {
		if e := nums[i]; e == val {
			valCount += 1

			target := l - valCount
			source := i

			nums[source], nums[target] = nums[target], nums[source]
		}
	}

	return len(nums) - valCount /* number of elements not equal to val */
}
