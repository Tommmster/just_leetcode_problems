package main

//You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//
//Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
//Return the maximum amount of water a container can store.
//
//Notice that you may not slant the container.
func maxAreaNaive(height []int) int {
	//naive

	l := len(height)

	if l <= 1 {
		return 0
	}

	maxArea := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			b := j - i
			h := min(height[i], height[j])

			maxArea = max(b*h, maxArea)
		}
	}
	return maxArea
}

// maxArea
// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.
func maxArea(height []int) int {
	//naive

	l := len(height)

	if l <= 1 {
		return 0
	}

	s, e := 0, l-1

	maxArea := 0
	for s < e {
		b := e - s
		h := min(height[s], height[e])

		maxArea = max(maxArea, b*h)

		if height[s] <= height[e] {
			s++
		} else {
			e--
		}
	}

	return maxArea
}
