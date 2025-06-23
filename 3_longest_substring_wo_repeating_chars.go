package main

// Given a string s, find the length of the longest without duplicate characters.
func lengthOfLongestSubstring(s string) int {
	var (
		l = len(s)
	)

	if l <= 1 {
		return l
	}

	var (
		start, end = 0, 1
		seen       = map[string]int{string(s[0]): 0}
		maxw       = end - start
	)

	for i := 1; i < l; i++ {
		v := string(s[i])

		if previous, ok := seen[v]; ok { // v was already in the window
			// Was this window better than the one before ?
			maxw = max(end-start, maxw)

			// Where do we set the new starting point? i points to a character that's repeated
			// start points to the beginning  of the current window
			start = previous + 1

			// mark the next character as the end. The new window spans from the
			// character following the previous value to the value folowwing the current
			// one
			end = i + 1

			//it's a new window, map should contain it's values
			seen = map[string]int{}
			for j := start; j < end; j++ {
				vv := string(s[j])
				seen[vv] = j
			}
		} else { // e is brand new, register it
			seen[v] = i
			end += 1
		}
	}

	return max(maxw, end-start)
}
