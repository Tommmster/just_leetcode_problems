package main

import (
	"cmp"
	"fmt"
	"reflect"
	"slices"
	"testing"
)

// 55 jump game
func TestCanJump(t *testing.T) {
	tt := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Simple success path",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "Stuck at zero",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "Single element",
			nums:     []int{0},
			expected: true, // already at the end
		},
		{
			name:     "Zero at start, can't move",
			nums:     []int{0, 1},
			expected: false,
		},
		{
			name:     "Just enough to reach end",
			nums:     []int{2, 0, 0},
			expected: true,
		},
		{
			name:     "pick up momentum",
			nums:     []int{2, 5, 0, 0},
			expected: true,
		},
		{
			name:     "just enough",
			nums:     []int{3, 0, 0, 0},
			expected: true,
		},
		{
			name:     "out of fuel in the middle",
			nums:     []int{1, 0, 2},
			expected: false,
		},
		{
			name:     "early true",
			nums:     []int{3, 0, 8, 2, 0, 0, 1},
			expected: true,
		},
		{
			name:     "improve the fuel situation when possible",
			nums:     []int{1, 1, 2, 2, 0, 1, 1},
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if result := canJump(tc.nums); result != tc.expected {
				t.Errorf("canJump(%v) = %v; expected %v", tc.nums, result, tc.expected)
			}
		})
	}
}

// 27 remove element
func TestRemoveElement(t *testing.T) {
	t.Run("return the number of elements not equal to val", func(t *testing.T) {
		x := removeElement([]int{1, 2, 3}, 2)

		if x != 2 {
			t.Errorf("[%s] incorrect number : %d", t.Name(), x)
		}
	})

	t.Run("when needle not found, nums is unchanged", func(t *testing.T) {
		var (
			nums = []int{1, 2, 3, 2, 3, 1}
			val  = 25
		)
		x := removeElement(nums, val)
		if x != 6 {
			t.Errorf("incorrect number :%d", x)
		}

		accepted := nums[:x]
		if !reflect.DeepEqual(nums, accepted) {
			t.Errorf("%s %+v Faulty array, expected %+v", t.Name(), accepted, nums)
		}

	})

	t.Run("multiple occurrences", func(t *testing.T) {
		var (
			nums = []int{1, 2, 3, 2, 3, 1}
			val  = 2
		)

		x := removeElement(nums, val)
		if x != 4 {
			t.Errorf("incorrect number :%d", x)
		}

		accepted := nums[:x]

		slices.SortStableFunc(accepted, func(a, b int) int {
			return cmp.Compare(a, b)
		})

		if !reflect.DeepEqual(accepted, []int{1, 1, 3, 3}) {
			t.Errorf("acc: %+v", accepted)
		}
	})

	t.Run("empty array should return 0", func(t *testing.T) {
		var (
			nums = []int{}
			val  = 2
		)

		x := removeElement(nums, val)
		if x != 0 {
			t.FailNow()
		}

	})
}

// 26
func TestRemoveDuplicates(t *testing.T) {
	t.Run("unique elements remain unique", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := removeDuplicates(nums)

		if k != 3 {
			t.Errorf("unexpected count: %d", k)
		}

		if !reflect.DeepEqual(nums[:k], nums) {
			t.Fatal()
		}
	})

	t.Run("Single unique element", func(t *testing.T) {
		var (
			nums     = []int{1, 1, 1, 1}
			expected = []int{1}
		)
		k := removeDuplicates(nums)

		if sorted := nums[:k]; !reflect.DeepEqual(sorted, expected) {
			t.Errorf("unexpected sorted result: %+v", sorted)
		}
	})

	t.Run("Remove duplicate elements", func(t *testing.T) {
		tt := []struct {
			nums     []int
			expected []int
		}{
			{nums: []int{}, expected: []int{}},
			{nums: []int{1}, expected: []int{1}},
			{nums: []int{1, 1, 1, 2}, expected: []int{1, 2}},
			{nums: []int{1, 1, 1, 2, 2, 3}, expected: []int{1, 2, 3}},
		}

		for i, tc := range tt {
			var (
				nums     = tc.nums
				expected = tc.expected
			)
			k := removeDuplicates(nums)

			if sorted := nums[:k]; !reflect.DeepEqual(sorted, expected) {
				t.Errorf("Test %s, Case %d: unexpected sorted result: %+v", t.Name(), i, sorted)
			}
		}
	})
}

//80
func TestRemoveDuplicates2(t *testing.T) {
	t.Run("unique elements remain unique", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := removeDuplicates2(nums)

		if k != 3 {
			t.Errorf("unexpected count: %d", k)
		}

		if !reflect.DeepEqual(nums[:k], nums) {
			t.Fatal()
		}
	})

	t.Run("Single unique element", func(t *testing.T) {
		var (
			nums     = []int{1, 1, 1, 1}
			expected = []int{1, 1}
		)
		k := removeDuplicates2(nums)

		if sorted := nums[:k]; !reflect.DeepEqual(sorted, expected) {
			t.Errorf("unexpected sorted result: %+v", sorted)
		}
	})

	t.Run("Remove duplicate elements", func(t *testing.T) {
		tt := []struct {
			nums     []int
			expected []int
		}{
			{nums: []int{}, expected: []int{}},
			{nums: []int{1}, expected: []int{1}},
			{nums: []int{1, 1, 1, 2}, expected: []int{1, 1, 2}},
			{nums: []int{1, 1, 1, 2, 2, 3}, expected: []int{1, 1, 2, 2, 3}},
			{nums: []int{1, 2, 2, 2, 3}, expected: []int{1, 2, 2, 3}},
			{nums: []int{1, 2, 2, 2, 2, 3, 4, 5}, expected: []int{1, 2, 2, 3, 4, 5}},
			{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, expected: []int{0, 0, 1, 1, 2, 3, 3}},
		}

		for i, tc := range tt {
			var (
				nums     = tc.nums
				expected = tc.expected
			)
			k := removeDuplicates2(nums)

			if sorted := nums[:k]; !reflect.DeepEqual(sorted, expected) {
				t.Errorf("Test %s, Case %d: unexpected sorted result: %+v, expected %+v", t.Name(), i, sorted, expected)
			}
		}
	})
}

// 88. Merge sorted arrays
func TestMerge(t *testing.T) {

	t.Run("a is empty", func(t *testing.T) {
		a := []int{0, 0, 0}
		b := []int{1, 2, 3}
		merge(a, b, 0, 3)

		if !reflect.DeepEqual(a, b) {
			t.Fatal("expected b")
		}
	})

	t.Run("b is empty", func(t *testing.T) {
		a := []int{6, 7, 8}
		merge(a, []int{}, 3, 0)
		if !reflect.DeepEqual(a, a) {
			t.Fatal("expected a")
		}
	})

	t.Run("a is lower than b", func(t *testing.T) {
		a := []int{1, 2, 3, 0, 0, 0}
		b := []int{4, 5, 6}
		merge(a, b, 3, 3)
		if !reflect.DeepEqual(a, []int{1, 2, 3, 4, 5, 6}) {
			t.Errorf("unexpected : %+v", a)
		}
	})

	t.Run("b is lower than a", func(t *testing.T) {
		a := []int{4, 5, 6, 0, 0, 0}
		b := []int{1, 2, 3}
		merge(a, b, 3, 3)
		if !reflect.DeepEqual(a, []int{1, 2, 3, 4, 5, 6}) {
			t.Errorf("unexpected : %+v", a)
		}
	})
	t.Run("Can merge both arrays of the same size", func(t *testing.T) {
		a := []int{2, 4, 6, 0, 0, 0}
		b := []int{1, 3, 5}
		merge(a, b, 3, 3)
		if !reflect.DeepEqual(a, []int{1, 2, 3, 4, 5, 6}) {
			t.Errorf("unexpected : %+v", a)
		}
	})
	t.Run("Can merge when b is longer", func(t *testing.T) {
		var (
			a        = []int{2, 4, 6, 0, 0, 0, 0, 0, 0}
			b        = []int{1, 3, 5, 7, 9, 11}
			expected = []int{1, 2, 3, 4, 5, 6, 7, 9, 11}
		)

		merge(a, b, 3, 6)

		if !reflect.DeepEqual(a, expected) {
			t.Errorf("unexpected : %+v", a)
		}
	})
	t.Run("Can merge when a is longer", func(t *testing.T) {
		var (
			a        = []int{2, 4, 6, 0, 0}
			b        = []int{1, 3}
			expected = []int{1, 2, 3, 4, 6}
		)

		merge(a, b, 3, 2)

		if !reflect.DeepEqual(a, expected) {
			t.Errorf("unexpected : %+v", a)
		}
	})

}

// 150: Rotate array
func TestRotateArray(t *testing.T) {

	getNums := func() []int {
		return []int{1, 2, 3, 4, 5}
	}
	var (
		singleRotation = []int{5, 1, 2, 3, 4}
		doubleRotation = []int{4, 5, 1, 2, 3}
		tripleRotation = []int{3, 4, 5, 1, 2}
	)

	t.Run("Single element", func(t *testing.T) {
		rotate([]int{1}, 1)

	})
	t.Run("Rotation by 0", func(t *testing.T) {
		nums := getNums()
		rotate(nums, 0)
		if !reflect.DeepEqual(nums, getNums()) {
			t.Fatalf("expected no change for k=0, got %+v", nums)
		}
	})
	t.Run("Rotation greater than array length", func(t *testing.T) {
		nums := getNums()
		rotate(nums, 7) // 7 % 5 == 2
		if !reflect.DeepEqual(nums, doubleRotation) {
			t.Fatalf("expected %+v, got %+v", doubleRotation, nums)
		}
	})

	t.Run("test rotations", func(t *testing.T) {
		tt := []struct {
			k        int
			expected []int
		}{
			{k: 1, expected: singleRotation},
			{k: 2, expected: doubleRotation},
			{k: 3, expected: tripleRotation},
		}

		for _, tc := range tt {
			var (
				nums = getNums()
				k    = tc.k
			)
			if rotate(nums, k); !reflect.DeepEqual(nums, tc.expected) {
				t.Errorf("unexpected rotation(%d): Expected: %+v, Got: %+v \n", k, tc.expected, nums)
			}

		}
	})
}

func TestMaxArea(t *testing.T) {

	tt := []struct {
		h []int
		a int
	}{
		{h: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, a: 49},
		{h: []int{1, 2}, a: 1},
	}

	for _, tc := range tt {

		if ma := maxArea(tc.h); ma != tc.a {
			t.Errorf("Unexpected area: Expected %d, Got %d\n", tc.a, ma)
		}
	}
}

// 8: Atoi
func TestAtoi(t *testing.T) {
	tt := []struct {
		s        string
		expected int
	}{
		{s: "42", expected: 42},
		{s: "+42", expected: 42},
		{s: "+0022", expected: 22},
		{s: "-58", expected: -58},
		{s: "  -17", expected: -17},
		{s: "ABCDEF", expected: 0},
		{s: "12F34", expected: 12},
		{s: "0-1", expected: 0},
		{s: "words and 987", expected: 0},
		{s: "-91283472332", expected: -2147483648},
		{s: "-2147483648", expected: -2147483648},
		{s: "+-12", expected: 0},
		{s: "-+12", expected: 0},
		{s: "21474836460", expected: 2147483647},
		{s: "  -0012a42", expected: -12},
		{s: "  +  413", expected: 0},
	}

	for i, tc := range tt {
		if n := myAtoi(tc.s); n != tc.expected {
			t.Errorf("Test %d (%s) failed: Expected %d, got %d\n", i, tc.s, tc.expected, n)
		}
	}
}

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "no flips needed",
			nums:     []int{1, 1, 1, 1, 1},
			k:        2,
			expected: 5,
		},
		{
			name:     "flip one zero in the middle",
			nums:     []int{1, 0, 1, 1, 0},
			k:        1,
			expected: 4,
		},
		{
			name:     "flip two zeros",
			nums:     []int{0, 0, 1, 1, 1, 0, 0},
			k:        2,
			expected: 5,
		},
		{
			name:     "not enough flips to merge groups",
			nums:     []int{1, 0, 1, 0, 1},
			k:        1,
			expected: 3,
		},
		{
			name:     "all zeros, flip all",
			nums:     []int{0, 0, 0},
			k:        3,
			expected: 3,
		},
		{
			name:     "all zeros, flip some",
			nums:     []int{0, 0, 0, 0},
			k:        2,
			expected: 2,
		},
		{
			name:     "no flips allowed",
			nums:     []int{1, 0, 1, 1, 0, 1},
			k:        0,
			expected: 2,
		},
		{
			name:     "empty input",
			nums:     []int{},
			k:        1,
			expected: 0,
		},
		{
			name:     "move window starting with zeroes",
			nums:     []int{0, 0 /**/, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1 /**/, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestOnes(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("%d/ longestOnes(%v, %d) = %d; want %d", i, tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func TestLongestUniqueSubstring(t *testing.T) {
	tt := []struct {
		s        string
		expected int
	}{
		{
			s:        "abcabcbb",
			expected: 3, // "abc"
		},
		{
			s:        "bbbbb",
			expected: 1, // "b"
		},
		{
			s:        "pwwkew",
			expected: 3, // "wke"
		},
		{
			s:        "",
			expected: 0, // empty string
		},
		{
			s:        " ",
			expected: 1, // single space
		},
		{
			s:        "au",
			expected: 2, // "au"
		},
		{
			s:        "dvdf",
			expected: 3, // "vdf"
		},
		{
			s:        "abba",
			expected: 2, // "ab" or "ba"
		},
		{
			s:        "abcdef",
			expected: 6, // whole string is unique
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			t.Parallel()

			if got := lengthOfLongestSubstring(tc.s); got != tc.expected {
				t.Errorf("failed: input=%s, got=%d, expected=%d", tc.s, got, tc.expected)
			}
		})
	}
}

func TestProductExceptSelf(t *testing.T) {
	tt := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			nums:     []int{0, 1, 2, 3},
			expected: []int{6, 0, 0, 0}, // only one zero
		},
		{
			nums:     []int{0, 0, 1, 2},
			expected: []int{0, 0, 0, 0}, // two zeros, all outputs are 0
		},
		{
			nums:     []int{5},
			expected: []int{1}, // edge case: single element
		},
		{
			nums:     []int{2, 3},
			expected: []int{3, 2},
		},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0}, // zero in center
		},
	}

	for i, tc := range tt {
		t.Skip()

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := productExceptSelf(tc.nums)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("productExceptSelf(%v) = %v; expected %v", tc.nums, got, tc.expected)
			}
		})
	}
}

//169
func TestMajorityElement(t *testing.T) {

	tests := []struct {
		nums     []int
		expected int
	}{
		// passing cases
		{[]int{6, 5, 5}, 5},
		{[]int{5, 5, 6}, 5},
		{[]int{2, 2, 1}, 2},
		{[]int{1, 1, 1, 2, 3}, 1},
		{[]int{9, 9, 9, 1, 1}, 9},

		// failing case for YOUR CURRENT IMPLEMENTATION:
		{[]int{3, 2, 3, 2, 2}, 2},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := majorityElement(tt.nums)
			if got != tt.expected {
				t.Errorf("majorityElement(%v) = %d, want %d",
					tt.nums, got, tt.expected)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic example",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "two elements",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "unsorted array",
			nums:   []int{1, 4, 5, 6},
			target: 10,
			want:   []int{1, 3},
		},
		{
			name:   "negatives included",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "mixed positives and negatives",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
		{
			name:   "zeros",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
		{
			name:   "large numbers",
			nums:   []int{1000000, 3, 999997},
			target: 1000000,
			want:   []int{1, 2},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestIsHappy(t *testing.T) {
	tests := []struct {
		input  int
		expect bool
	}{
		{1, true},
		{7, true},
		{10, true},
		{19, true},
		{2, false},
		{3, false},
		{4, false},
		{20, false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isHappy(tt.input)
			if got != tt.expect {
				t.Errorf("isHappy(%d) = %v, expected %v", tt.input, got, tt.expect)
			}
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input  []string
		expect string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interview", "internet", "internal"}, "inter"},
		{[]string{"a", "a", "a"}, "a"},
		{[]string{"", "abc", "abcd"}, ""},
		{[]string{"single"}, "single"},
		{[]string{}, ""},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := longestCommonPrefix(tt.input)
			if got != tt.expect {
				t.Errorf("longestCommonPrefix(%v) = %v, expected %v", tt.input, got, tt.expect)
			}
		})
	}
}

func TestLowercaseTrieLCP(t *testing.T) {
	tests := []struct {
		words  []string
		expect string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interview", "internet", "internal"}, "inter"},
		{[]string{"a", "a", "a"}, "a"},
		{[]string{"", "abc", "abcd"}, ""},
		{[]string{"single"}, "single"},
		{[]string{}, ""},
		{[]string{"", ""}, ""},
		{[]string{"aa", "aa"}, "aa"},
		{[]string{"ab", "a"}, "a"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			trie := MakeLowercaseTrie()
			for _, w := range tt.words {
				_ = trie.Put(w)
			}
			got := trie.LongestCommonPrefix()
			if got != tt.expect {
				t.Errorf("LCP(%v) = [%v], expected [%v]", tt.words, got, tt.expect)
			}
		})
	}
}

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},

		{"", "anything", true},
		{"a", "", false},
		{"", "", true},

		{"a", "a", true},
		{"a", "b", false},

		{"aaa", "aaaaaa", true},
		{"aaaa", "aaa", false},

		{"bca", "abc", false},

		{"áç", "xáyçz", true},
		{"çá", "xáyçz", false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("isSubsequence(%q, %q) = %v, want %v",
					tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		// Basic cases
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"abc", "", 0},
		{"", "a", -1},

		// Single char
		{"a", "a", 0},
		{"a", "b", -1},

		// Overlapping
		{"aaa", "aa", 0},

		// Required: mississippi cases
		{"mississippi", "issip", 4},
		{"mississippi", "issi", 1},
		{"mississippi", "mississippia", -1},
		{"mississippi", "pi", 9},
		{"mississippi", "p", 8},

		// Additional robustness
		{"abcde", "e", 4},
		{"abcde", "de", 3},
		{"abcabcabcd", "abcabcd", 3},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := strStr(tc.haystack, tc.needle)
			if got != tc.want {
				t.Errorf("strStr(%q, %q) = %d, want %d",
					tc.haystack, tc.needle, got, tc.want)
			}
		})
	}
}

func TestMostCommon(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 2, 3}, 2},
		{[]int{5}, 5},
		{[]int{7, 7, 7}, 7},
		{[]int{1, 2, 3}, 0},
		{[]int{4, 4, 5, 5}, 0},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := MostCommon(tt.input)
			if tt.expected != 0 && got != tt.expected {
				t.Fatalf("got %d, expected %d", got, tt.expected)
			}
		})
	}
}

func TestRLE(t *testing.T) {
	t.Run("empty input should return empty output", func(t *testing.T) {
		if rleLike("") != "" {
			t.Fatal("should be empty")
		}
	})
	t.Run("encode", func(t *testing.T) {
		if encoded := rleLike("aaabbc"); encoded != "a3b2c1" {
			t.Errorf("mismatch: Got %s", encoded)
		}
	})

	t.Run("keep the shortest", func(t *testing.T) {
		if encoded := rleLike("abc"); encoded != "abc" {
			t.Errorf("mismatch: Got %s", encoded)
		}
	})
}
