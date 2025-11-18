package main

// You are given an integer array nums.
// You are initially positioned at the array's first index, and each element in the array represents
// your maximum jump length at that position.
// Return true if you can reach the last index, or false otherwise.
func canJump(a []int) bool {
	if len(a) <= 1 {
		return true
	}
	if a[0] == 0 {
		return false
	}

	var (
		last = len(a) - 1
		at   = 0
		d    = a[at] // how far can we move from the current point
	)

	// For each element of the array, the index indicates how many steps from
	// the start we've taken. The element at index i will tell how much fuel
	// we can gain from there.
	for i, e := range a[:last] {
		if jumps_to := i + e; jumps_to >= last { // can we make an early call ? The boost at this point would be enough
			return true
		}

		used := (i - at) // how many spaces have we moved with the current boost
		remaining := d - used

		if remaining == 0 { // out of fuel
			if e == 0 { // cant charge
				return false
			}
			// refill
			at = i
			d = e
		} else if e > remaining {
			// refill
			at = i
			d = e
		}
	}

	return true
}
