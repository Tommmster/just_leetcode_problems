package main

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
func rotate(nums []int, k int) {
	if k <= 0 {
		return
	}

	k = k % len(nums)

	// [1,2,3,4,5,6,7] len(.) = 7
	// [5,6,7,1,2,3,4] k = 3 :
	// Take the last 3 elements and place them at the front
	// (last three elements have indices 4,5,6 respectively)

	doReverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	doReverse(0, len(nums)-1)
	doReverse(0, k-1)
	doReverse(k, len(nums)-1)
}
