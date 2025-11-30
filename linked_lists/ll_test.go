package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLinkedListLength(t *testing.T) {
	tt := []struct {
		vals []int
		l    int
	}{
		{
			vals: []int{},
			l:    0,
		},
		{
			vals: []int{1},
			l:    1,
		},
		{
			vals: []int{1, 2, 3, 4, 5, 6},
			l:    6,
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			head := listFromSlice(tc.vals)
			if got := length(head); got != tc.l {
				t.Fatalf("[%s] Failed: Expected %d, got %d", t.Name(), tc.l, got)
			}
		})
	}
}

func TestLinkedListSplit(t *testing.T) {
	tt := []struct {
		val []int
		l   int
		el  []int // expected low half
		eh  []int // expected high half
	}{
		{
			val: []int{1, 2, 3, 4, 5},
			l:   2,
			el:  []int{1, 2},
			eh:  []int{3, 4, 5},
		},
		{
			val: []int{1, 2, 3, 4, 5},
			l:   10,
			el:  []int{1, 2, 3, 4, 5},
			eh:  []int{},
		},
		{
			val: []int{1, 2, 3, 4, 5},
			l:   0,
			el:  []int{},
			eh:  []int{1, 2, 3, 4, 5},
		},
		{
			val: []int{},
			l:   5,
			el:  []int{},
			eh:  []int{},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			head := listFromSlice(tc.val)
			gl, gh := split(head, tc.l)

			if got := sliceFromList(gl); !reflect.DeepEqual(tc.el, got) {
				t.Fatalf("Expected %v got %v", tc.el, got)
			}

			if len(tc.eh) == 0 && gh != nil {
				t.Fatal("Unexpected second half")
			}

			if got := sliceFromList(gh); !reflect.DeepEqual(tc.eh, got) {
				t.Fatalf("Expected %v got %v", tc.eh, got)
			}
		})
	}
}

func TestSortLinkedList(t *testing.T) {
	tt := []struct {
		vs []int
		ex []int
	}{
		{
			vs: []int{2, 1},
			ex: []int{1, 2},
		},
		{
			vs: []int{},
			ex: []int{},
		},
		{
			vs: []int{1},
			ex: []int{1},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			h := listFromSlice(tc.vs)
			s := SortList(h)
			if got := sliceFromList(s); !reflect.DeepEqual(tc.ex, got) {
				t.Fatalf("Expected %v, got %v\n", []int{1, 2}, got)
			}
		})
	}
}

func TestReverseLinkedList(t *testing.T) {
	mustSetupList := func(values ...int) *ListNode {
		return listFromSlice(values)
	}
	t.Run("Must reverse non trivial list", func(t *testing.T) {
		t.Parallel()
		head := mustSetupList(1, 2, 3, 4, 5)
		if head.Val != 1 {
			t.Fatalf("Wrong head, expected %d got %d\n", 1, head.Val)
		}

		head = reverseList(head)
		if head.Val != 5 {
			t.Fatalf("Wrong head, expected %d got %d\n", 5, head.Val)
		}
	})
	t.Run("Must reverse single element list", func(t *testing.T) {
		t.Parallel()
		head := mustSetupList(1)
		if head.Val != 1 {
			t.Fatalf("Wrong head, expected %d got %d\n", 1, head.Val)
		}

		head = reverseList(head)
		if head.Val != 1 {
			t.Fatalf("Wrong head, expected %d got %d\n", 1, head.Val)
		}
	})

	t.Run("Empty list returns itself", func(t *testing.T) {
		h := reverseList(nil)
		if h != nil {
			t.Fatal("Expected nil list")
		}
	})
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		list      []int
		deleteIdx int
		want      []int
	}{
		// Delete middle
		{[]int{1, 2, 3, 4}, 1, []int{1, 3, 4}}, // delete "2"
		{[]int{4, 5, 1, 9}, 1, []int{4, 1, 9}}, // delete "5"

		// Delete near the end (but not last)
		{[]int{10, 20, 30, 40}, 2, []int{10, 20, 40}}, // delete "30"

		// Delete head.next
		{[]int{7, 8, 9}, 0, []int{8, 9}}, // delete "7" (but not truly head since LC gives a pointer to node, not head)

		// Multiple nodes
		{[]int{1, 3, 5, 7, 9}, 2, []int{1, 3, 7, 9}}, // delete "5"
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			head := listFromSlice(tt.list)

			node := nodeAt(head, tt.deleteIdx)
			if node == nil || node.Next == nil {
				t.Fatalf("invalid test: node at index %d is nil or last", tt.deleteIdx)
			}

			deleteNode(node)

			got := sliceFromList(head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("after deleting index %d: got %v, want %v", tt.deleteIdx, got, tt.want)
			}
		})
	}
}

func TestRotateRight(t *testing.T) {
	tt := []struct {
		values []int
		k      int
		want   []int
	}{
		{
			values: []int{5},
			k:      1,
			want:   []int{5},
		},
		{
			values: []int{10, 11, 12},
			k:      0,
			want:   []int{10, 11, 12},
		},
		{
			values: []int{1, 2, 3, 4, 5},
			k:      1,
			want:   []int{5, 1, 2, 3, 4},
		},
		{
			values: []int{1, 2, 3, 4, 5},
			k:      2,
			want:   []int{4, 5, 1, 2, 3},
		},
		{
			values: []int{0, 1, 2},
			k:      4, // full rotation → same list
			want:   []int{2, 0, 1},
		},
		{
			values: []int{1, 2},
			k:      2,
			want:   []int{1, 2},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			head := listFromSlice(tc.values)

			got := rotateRight(head, tc.k)
			if expected := sliceFromList(got); !reflect.DeepEqual(expected, tc.want) {
				t.Fatalf("Wanted %v, got %v\n", tc.want, expected)
			}
		})
	}
}
