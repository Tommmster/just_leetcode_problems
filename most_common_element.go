package main

// returns the most common element
// e.g nums = [1, 1, 1, 2, 2, 3] -> 1
func MostCommon(nums []int) int {

	// Keep track of the first occurrence of each element
	var m = make(map[int]int)
	for _, e := range nums {
		if _, ok := m[e]; !ok {
			l := len(m) // current number of mappings
			m[e] = l
		}
	}

	// Count the occurrences for each element
	var uniqueCounts = make([]int, len(m))
	for _, e := range nums {
		i := m[e]
		uniqueCounts[i]++
	}

	// Now we need to find the max in the occurrences vector
	// This vector has as many elements as unique numbers in num
	var (
		max = uniqueCounts[0]
		at  = 0
	)

	for i := 1; i < len(uniqueCounts); i++ {
		if uniqueCounts[i] > max {
			at = i
			max = uniqueCounts[i]
		}
	}

	// Reverse-lookup the map to find which element we're dealing with
	for k, v := range m {
		if v == at {
			return k
		}
	}

	return 0
}
