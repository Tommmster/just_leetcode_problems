package main

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
func majorityElement(nums []int) int {
	// Boyer - Moore, Can use it if there's a majority element. Can't otherwise
	candidate := 0
	count := 0

	for _, e := range nums {
		if count == 0 {
			candidate = e
		}

		if e == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
