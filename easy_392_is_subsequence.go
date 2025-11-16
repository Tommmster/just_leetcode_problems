package main

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none)
// of the characters without disturbing the relative positions of the remaining characters.
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//
// Example 1:
//
// Input: s = "abc", t = "ahbgdc"
// Output: true
//
// Example 2:
//
// Input: s = "axc", t = "ahbgdc"
// Output: false

func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}
	var (
		got  = 0
		need = len(s)
		want = s[got]
	)

	for i := 0; i < len(t); i++ {
		e := t[i]
		if e == want {
			got++
			if got == need {
				return true
			}

			want = s[got]
		}

	}

	return false
}
