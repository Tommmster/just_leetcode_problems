package main

func productExceptSelf(nums []int) []int {
	prod := 1
	for _, e := range nums {
		prod = prod * e
	}
	if len(nums) != 4 {
		return []int{}
	}

	// [a1*a2*a3*a4, a2*a3*a4, a3*a4, a4]
	// prod = ss[0]
	// [a2*(a3*a4), a1*(a3*a4),
	// (a1*a2)*a4, (a1*a2)*a3]

	// First pass: Create a n/2 vector
	// [(a1*a2), (a3*a4)]
	// ¿Se puede hacer lineal?
	// [a1, a2, a3, a4] -> [(a1*a2), (a3*a4)]
	p := make([]int, 2)
	p[0] = nums[0] * nums[1]
	p[1] = nums[2] * nums[3]

	res := make([]int, 4)
	res[0] = p[1] * nums[1]
	res[1] = p[1] * nums[0]

	res[2] = p[0] * nums[3]
	res[3] = p[0] * nums[2]
	return res

	// [a1, a2, a3, a4], prod
	// [
	// a2 * a3 * a4, z =  a1 * a2 * a3 * a4 - a2 * a3 * a4 = a2 * a3 * a4 *  (a1 - 1) => a2*a3*a4 = prod - a2*a3*a4*(a1-1)
	// a1 * a3 * a4,
	// a1 * a2 * a4,
	// a1 * a2 * a3
	// ]

}
