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

		previous, ok := seen[v]
		if !ok || previous < start {
			seen[v] = i
			end += 1
		} else {
			// Was this window better than the one before ?
			maxw = max(end-start, maxw)

			// Where do we set the new starting point? i points to a character that's repeated
			// start points to the beginning  of the current window
			start = previous + 1

			// mark the next character as the end. The new window spans from the
			// character following the previous value to the value folowwing the current
			// one
			end = i + 1

			// Update the position of the relevant values
			atStart := string(s[start])
			seen[atStart] = start
			seen[v] = i
		}
	}

	return max(maxw, end-start)
}
