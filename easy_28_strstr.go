package main

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	} else if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0 // start at the beginning of the needle

		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i
		}
	}

	return -1
}
