package main

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
func majorityElement(nums []int) int {
	var (
		counts  = make([]int, len(nums))
		firstAt = make(map[int]int)
		at      = 0
		max     = 0
	)

	// Queremos contar cuantas veces aparece cada elemento. En counts tenemos los contadores
	for i, e := range nums {
		seenAt, ok := firstAt[e]
		if !ok {
			seenAt = i
			firstAt[e] = seenAt
		}
		counts[seenAt] += 1

		if c := counts[seenAt]; c > max {
			at = seenAt
			max = c
		}
	}

	return nums[at]
}
