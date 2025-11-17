package main

import (
	"errors"
)

// returns the most common element
// e.g nums = [1, 1, 1, 2, 2, 3] -> 1
func MostCommon(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("empty list")
	}

	// Keep a counter per element
	var m = make(map[int]int)
	for _, e := range nums {
		m[e]++
	}

	var (
		max = 0
		at  = 0
	)

	for k, v := range m {
		if v > max {
			max = v
			at = k
		}
	}

	return at, nil
}

// Given a list of numbers, return the index of the most common element.
// If the list is empty, return -1
// If there is more than one, return the first one
func MostCommonIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	// Use this struct as tuples are not an option.
	// Should be used as an immutable object.
	type vv struct {
		firstAt int
		count   int
	}

	// a map of unique elements, pointing to their first occurrence and keeping track
	// of the count
	var u = make(map[int]vv)
	for i, e := range nums {
		if p, ok := u[e]; !ok {
			u[e] = vv{i, 1}
		} else {
			u[e] = vv{p.firstAt, p.count + 1}
		}
	}

	var (
		max = 0
		at  = 0
	)

	// Now traverse the map so we can find the earliest appearence of the maximal element
	for _, v := range u {
		if v.count > max {
			max = v.count
			at = v.firstAt // I Want to return the first one, how do I achieve that
		} else if v.count == max && v.firstAt < at {
			at = v.firstAt
		}
	}

	return at
}
