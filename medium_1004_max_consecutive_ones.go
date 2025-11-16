package main

// Given a binary array nums and an integer k,
// return the maximum number of consecutive 1's in the array
// if you can flip at most k 0's.
func longestOnes(v []int, k int) int {
	var (
		s, e = -1, -1
		l    = len(v)
	)

	// list is empty, 0 is the natural choice
	if l == 0 {
		return 0
	}

	// init stuff
	//List starts with a 1, track a window starting from the beginning
	if d := v[0]; d == 1 {
		s, e = 0, 1
	} else if k > 0 { // First element is a 0, flip it and track a window
		s, e = 0, 1
		k -= 1
	} else { // first element is 0, no flips allowed. Don't track a window, move on to the next element
		if l == 1 {
			return 0
		}
		// window length is 0
		s, e = 1, 1
	}

	// track the amount of ones in a window, considering flips
	maxones := s - e
	for i := 1; i < l; i++ {
		d := v[i]
		if d == 1 {
			e += 1
		} else if k > 0 {
			k -= 1
			e += 1
		} else { // out of flips
			// Found a 0 and we're out of flips, so this window is over. Get the window length
			// as start - end and compare it to the current max
			maxones = max(e-s, maxones)

			// start might be pointing to a 1 or a 0
			for v[s] != 0 && s < l { // skip past the ones
				s++
			}
			if s < l {
				s++ // skip past the earliest 0 from start
				e++ //move the end one further, the current 0 is flipped to a one
			}
		}
	}
	if s >= l {
		return maxones
	}

	return max(maxones, e-s)
}
