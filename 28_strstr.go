package main

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	} else if len(haystack) < len(needle) {
		return -1
	}

	var (
		vs = []int{} // starting points
	)

	// Look for points in the haystack where needle might begin
	for i := 0; i < len(haystack); i++ {
		e := haystack[i]
		t := needle[0]
		if e == t && i+len(needle) <= len(haystack) {
			vs = append(vs, i)
		}
	}

	// Now go through those starting points only
	for j := 0; j < len(vs); j++ {
		if check_hs(haystack, vs[j], vs[j]+len(needle), needle) {
			return vs[j]
		}
	}

	return -1
}

func check_hs(h string, s, e int, needle string) bool {
	j := 0
	for i := s; i < e; i++ {

		if h[i] != needle[j] {
			return false
		}
		j++

	}
	return true
}
