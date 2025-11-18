package main

import (
	"fmt"
	"reflect"
	"testing"
)

// This is an exercise on the sliding window, used to solve
// a number of problems (e.g finding the largest palindrome, etc)
// The exercise at hand is to find all the windows, specified by the start(inclusive)
// end end(exclusive) indexes
// wboundaries
type wboundaries struct {
	s, e int
}

//wb: Find the [start, end) boundaries of the window with at most k zero elements
func wb(v []int, k int) []wboundaries {
	var (
		s, e = -1, -1
		wbs  = []wboundaries{}
		l    = len(v)
	)

	if l == 0 {
		return wbs
	}

	// init stuff
	if d := v[0]; d == 1 {
		s, e = 0, 1
	} else if k > 0 {
		s, e = 0, 1
		k -= 1
	} else { // first element is 0, no flips allowed.
		if l == 1 {
			return wbs
		}
		s, e = 1, 1
	}

	for i := 1; i < l; i++ {
		d := v[i]
		if d == 1 {
			e += 1
		} else if k > 0 {
			k -= 1
			e += 1
		} else {
			// out of flips
			// append the current boundaries
			wbs = append(wbs, wboundaries{s, e})

			// start might be pointing to a 1 or a 0
			for v[s] != 0 && s < l { // skip past the ones
				s++
			}
			if s < l {
				s++ // skip past the earliest 0 from start
				e++ //move the end one further, the current 0 is flipped to a one
			}
		}
	}
	if s >= l {
		return wbs
	}

	return append(wbs, wboundaries{s, e})
}

func TestWindowBoundaries(t *testing.T) {
	tt := []struct {
		v        []int
		k        int
		expected []wboundaries
	}{
		{
			v:        []int{1, 0, 1, 0},
			k:        1,
			expected: []wboundaries{{0, 3}, {2, 4}},
		},
		{
			v:        []int{1, 1, 1, 1},
			k:        1,
			expected: []wboundaries{{0, 4}}, // no flips needed
		},
		{
			v:        []int{0, 0, 0, 0},
			k:        2,
			expected: []wboundaries{{0, 2}, {1, 3}, {2, 4}},
		},
		{
			v:        []int{1, 0, 1, 1, 0, 1, 1, 1},
			k:        1,
			expected: []wboundaries{{0, 4}, {2, 8}},
		},
		{
			v:        []int{},
			k:        2,
			expected: []wboundaries{},
		},
		{
			v:        []int{0, 1, 0, 1, 0},
			k:        0,
			expected: []wboundaries{{1, 2}, {3, 4}}, // only natural 1s
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			wbs := wb(tc.v, tc.k)

			if !reflect.DeepEqual(tc.expected, wbs) {
				t.Errorf("%d/ failed, (%+v, %d) = (%+v), expected (%+v)", i, tc.v, tc.k, wbs, tc.expected)
			}
		})
	}
}
