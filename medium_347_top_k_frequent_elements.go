package main

func topKFrequent(nums []int, k int) []int {

	var counter = make(map[int]int)
	var max = 0
	for _, e := range nums {
		v := counter[e]
		v++
		counter[e] = v
		if v > max {
			max = v
		}
	}

	var buckets = make([][]int, max)
	for k, v := range counter {
		buckets[v-1] = append(buckets[v-1], k)
	}

	var (
		// k-top results
		res       = []int{}
		count int = k
	)

	for j := len(buckets) - 1; j >= 0 && count > 0; j-- {
		if len(buckets[j]) == 0 {
			continue
		}

		l := len(buckets[j])
		if l >= k {
			return buckets[j][:k]
		} else { // hay menos valores de los que quiero. Copio todos
			res = append(res, buckets[j]...)
			count -= l
		}
	}

	return res
}
