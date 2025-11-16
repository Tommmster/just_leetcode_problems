package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	l := len(strs[0])
	for _, e := range strs[1:] {
		l = min(l, len(e))
	}

	var found bool
	var longestPrefix string
	for i := 1; i <= l; i++ {
		candidate := strs[0][:i]
		found = true

		for _, e := range strs[1:] {
			if e[:i] != candidate {
				found = false
				break
			}
		}

		if found {
			longestPrefix = candidate
		} else {
			break
		}
	}
	return longestPrefix
}
