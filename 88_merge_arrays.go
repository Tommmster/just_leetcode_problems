package main

func merge(a []int, b []int, m int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		copy(a, b)
	}

	pa := m - 1
	pb := n - 1
	target := m + n - 1

	for pa >= 0 && pb >= 0 {
		if a[pa] > b[pb] {
			a[target] = a[pa]
			pa--
		} else {
			a[target] = b[pb]
			pb--
		}
		target--
	}

	if pb >= 0 {
		copy(a[:pb+1], b[:pb+1])
	}
}
