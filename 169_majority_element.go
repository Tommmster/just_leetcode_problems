package main

import "fmt"

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
func majorityElement(nums []int) int {
	var (
		counts = make([]int, len(nums))
		xxx    = make(map[int]int)
	)

	// Queremos contar cuantas veces aparece cada elemento. En counts tenemos los contadores
	for i, e := range nums {
		seenAt, ok := xxx[e]
		if ok {
			counts[seenAt] += 1
		} else {
			xxx[e] = i
			counts[i] = 1
		}
	}

	var (
		max = counts[0]
		at  = 0
	)
	for i, e := range counts[1:] {
		if e > max {
			max = e
			at = i + 1
		}
	}
	fmt.Printf("%v max: %d maxAt: %d\n", counts, max, at)

	return nums[at]
}
