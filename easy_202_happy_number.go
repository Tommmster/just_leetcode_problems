package main

// Write an algorithm to determine if a number n is happy.
//
// A happy number is a number defined by the following process:
//
//     Starting with any positive integer, replace the number by the sum of the squares of its digits.
//     Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
//     Those numbers for which this process ends in 1 are happy.
//
// Return true if n is a happy number, and false if not.
// E.g 19 -> 1 + 81 = 82 -> 64 + 4 = 68 -> 36 + 64 = 100 -> 1

func isHappy(n int) bool {

	seen := make(map[int]struct{})

	return isHappyInner(n, seen)
}

func isHappyInner(n int, seen map[int]struct{}) bool {
	digits := []int{}

	if _, ok := seen[n]; ok {
		return false
	}

	seen[n] = struct{}{}

	for n > 0 {
		r := n % 10
		n = n / 10

		digits = append(digits, r)
	}
	if n > 0 {
		digits = append(digits, n)
	}

	// Add the squares
	sqsum := 0
	for _, e := range digits {
		sqsum += e * e
	}

	if sqsum == 1 {
		return true
	}
	return isHappyInner(sqsum, seen)
}
