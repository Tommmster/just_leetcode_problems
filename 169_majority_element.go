package main

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
func majorityElement(nums []int) int {
	var (
		firstAt = make(map[int]int)
		at      = 0
		max     = 0
	)

	// Queremos contar cuantas veces aparece cada elemento. En counts tenemos los contadores
	for i, e := range nums {
		seenAt, ok := firstAt[e]

		// Es la primea vez que vemos el numero e
		if !ok {
			// Lo marcamos como visto por primera vez en i
			seenAt = i
			firstAt[e] = seenAt

			// Usamos esta posicion como contador
			nums[seenAt] = 0
		}

		// Si ya habiamos visto el numero e, entonces seenAt apunta al contador
		nums[seenAt]++
		if nums[seenAt] > max {
			at = seenAt
			max = nums[seenAt] //max count
		}
	}

	for k, v := range firstAt {
		if v == at {
			return k
		}
	}

	panic("should not happen")
}
