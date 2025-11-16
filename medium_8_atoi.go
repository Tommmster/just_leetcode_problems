package main

// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
// The algorithm for myAtoi(string s) is as follows:
//
//     Whitespace: Ignore any leading whitespace (" ").
//     Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
//     Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
//     Rounding: If the integer is out of the 32-bit signed integer range [-2^31, 2^31 - 1], then round the integer to remain in the range. Specifically, integers less than -2^31 should be rounded to -2^31, and integers greater than 2^31 - 1 should be rounded to 2^31 - 1.
//
// Return the integer as the final result.
func myAtoi(s string) int {
	const (
		upperBound int64 = (1 << 31) - 1
		lowerBound int64 = -1 << 31
	)

	var (
		n          = 0
		startAt    = 0
		isNegative bool
		isPositive bool
	)

	allowedAt := func(i int) bool {
		if isNegative || isPositive {
			return (s[i] >= '0' && s[i] <= '9')
		}
		return (s[i] >= '0' && s[i] <= '9') || s[i] == '-' || s[i] == '+' || s[i] == ' '
	}
	for i := 0; i < len(s); i++ {

		if !allowedAt(i) {
			return 0
		}

		if s[i] >= '0' && s[i] <= '9' {
			break
		}
		if s[i] == '+' {
			isPositive = true
		} else if s[i] == '-' {
			isNegative = true
		}

		if s[i] == '-' {
			isNegative = true
		}

		startAt = startAt + 1
	}

	if startAt == len(s) {
		return 0
	}

	for i := startAt; i < len(s); i++ {
		if nonDigit := s[i] < '0' || s[i] > '9'; nonDigit {
			if isNegative {
				return -n
			}

			return n
		}

		n = n*10 + int(s[i]-'0')

		if isNegative && -int64(n) <= lowerBound {
			return int(lowerBound)
		}
		if int64(n) > upperBound {
			return int(upperBound)
		}
	}

	if isNegative {
		return -n
	}

	return n

}
