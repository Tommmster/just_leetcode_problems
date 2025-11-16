package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, e := range nums {
		v, ok := m[target-e]
		if ok && v != i {
			return []int{v, i}
		}

		m[e] = i
	}
	return []int{}
}
