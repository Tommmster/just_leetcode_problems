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
		seen       = map[byte]int{s[0]: 0}
		maxw       = end - start
	)

	for i := 1; i < l; i++ {
		v := s[i]
		if previous, ok := seen[v]; ok && previous >= start {
			// if the current character has already been seen,
			// the start is the character next to its previous appearence
			start = previous + 1
		}
		seen[v] = i
		end = i + 1 // the end of the window always moves forward
		maxw = max(end-start, maxw)
	}

	return maxw
}
