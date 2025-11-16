package main

func productExceptSelf(nums []int) []int {
	// Given nums: [a1, a2, a3, a4, a5]
	// Set up groups
	// [a2*(a3*a4*a5), a1*(a3*a4*a5),
	// a4*(a1*a2*a5), a3*(a1*a2*a5),
	// a1*a2*a3*a4]

	// First pass: Create a n/2 vector, where each element contains the product of half the numbers
	// [(a3*a4), (a1*a2)]
	var (
		l  = len(nums)
		lp = l / 2
		p  = make([]int, lp)
	)

	// [a1,
	//  a2 * a1,
	//  a3*a2*a1,
	//  a4*a3*a2*a1,
	//  a5*a3*a4*a2*a1
	// ]

	// [a5, a5*a4, a5*a4*a3, a5*a4*a3*a2, a5*a4*a3*a2*a1]
	// [a5*a4*a3*a2, a5*a4*a3, a5*a4, a4, a5]
	// Como generalizamos para n elementos
	for i := 0; i < len(nums); i++ {
		// i: 0,1,2,3,4

	}
	// Set up groups
	for i := range lp {
		pindex := (lp - 1) - i
		p[pindex] = nums[2*i] * nums[2*i+1]
	}

	// [a2*(a3*a4), a1*(a3*a4),
	// (a1*a2)*a4, (a1*a2)*a3]
	//Swap values
	for i := range lp {
		nums[2*i], nums[2*i+1] = nums[2*i+1], nums[2*i]
	}
	//Calculate in place
	for i := 0; i < len(p); i++ {
		nums[i] = nums[i] * p[0]         // first half
		nums[l/2+i] = nums[l/2+i] * p[1] // second half
	}

	return nums
}
