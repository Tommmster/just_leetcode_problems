package main

func productExceptSelf(nums []int) []int {
	prod := 1
	for _, e := range nums {
		prod = prod * e
	}
	if len(nums) != 4 {
		panic("not implemented")
	}

	// [a1*a2*a3*a4, a2*a3*a4, a3*a4, a4]
	// prod = ss[0]
	// [a2*(a3*a4), a1*(a3*a4),
	// (a1*a2)*a4, (a1*a2)*a3]

	// First pass: Create a n/2 vector
	// [(a1*a2), (a3*a4)]
	// ¿Se puede hacer lineal?
	// [a1, a2, a3, a4] -> [(a1*a2), (a3*a4)]

	l := len(nums)
	p := make([]int, l/2)
	lp := len(p)

	// [(a3*a4), (a1*a2)]
	for i := 0; i < lp; i++ {
		pindex := (lp - 1) - i
		p[pindex] = nums[2*i] * nums[2*i+1]
	}

	res := make([]int, l)
	for i := 0; i < len(p); i++ { // first half: 0, 1
		nindex := (l/2 - 1) - i
		res[i] = nums[nindex] * p[0]
	}

	for i := 0; i < len(p); i++ { // second half: 0,1
		nindex := (l - 1) - i // traverse the second half backwards: 3, 2
		rindex := (l / 2) + i // traverse the second half forwards: 2,3
		res[rindex] = nums[nindex] * p[1]
	}

	return res
}
