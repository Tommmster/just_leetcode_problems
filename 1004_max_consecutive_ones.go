package main

// Given a binary array nums and an integer k,
// return the maximum number of consecutive 1's in the array
// if you can flip at most k 0's.
func longestOnes(nums []int, k int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	n, r := 0, k
	current := 0

	for _, e := range nums {
		if e == 1 {
			current += e
		} else if r > 0 && current > 0 {
			r -= 1
			current += 1
		} else {
			var nn int
			if k > 0 { // reset the counter
				r = k - 1
				nn = 1
			}
			n, current = max(current, n), nn
		}
	}
	return max(current, n)
}
