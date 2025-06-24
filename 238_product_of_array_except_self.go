package main

func productExceptSelf(nums []int) []int {
	// [a1*a2*a3*a4, a2*a3*a4, a3*a4, a4]
	// prod = ss[0]
	// [a2*(a3*a4), a1*(a3*a4),
	// (a1*a2)*a4, (a1*a2)*a3]

	// First pass: Create a n/2 vector, where each element contains the product of half the numbers
	// [(a3*a4), (a1*a2)]
	var (
		l  = len(nums)
		lp = l / 2
		p  = make([]int, lp)
	)

	for i := 0; i < lp; i++ {
		pindex := (lp - 1) - i
		p[pindex] = nums[2*i] * nums[2*i+1]
	}

	// [a2*(a3*a4), a1*(a3*a4),
	// (a1*a2)*a4, (a1*a2)*a3]
	for i := range lp {
		nums[2*i], nums[2*i+1] = nums[2*i+1], nums[2*i]
	}

	for i := 0; i < len(p); i++ { // first half
		nums[i] = nums[i] * p[0]
	}

	for i := 0; i < len(p); i++ { // second half
		nums[l/2+i] = nums[l/2+i] * p[1]
	}

	return nums
}
