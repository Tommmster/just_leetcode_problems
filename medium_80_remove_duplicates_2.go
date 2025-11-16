package main

// Given an integer array nums sorted in non-decreasing order,
// remove some duplicates in-place such that each unique element appears at most twice.
// The relative order of the elements should be kept the same.
func removeDuplicates2(nums []int) int {

	const keep = 2

	if l := len(nums); l <= keep {
		return l
	}

	type seqState struct {
		current int
		count   int
	}

	state := seqState{
		current: nums[0],
		count:   1,
	}

	// The array invariant (it's sorted in a non-descending order) should be true
	// for elements between indices 0 and k (exclusive)
	k := 1

	for _, e := range nums[1:] {
		if e == state.current {
			if state.count < keep {
				k += 1
			}
			state.count += 1
			if state.count >= keep {
				nums[k-1] = e
			}
		} else {
			state = seqState{
				current: e,
				count:   1,
			}
			nums[k] = e
			k += 1
		}
	}

	return k
}
