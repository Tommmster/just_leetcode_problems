package containers

import "testing"

func TestHeap(t *testing.T) {
	setUp := func(maxSize int, values ...int) *Heap {
		return NewHeap(maxSize, values...)
	}

	t.Run("peek returns but doesn't remove", func(t *testing.T) {
		heap := setUp(5, 3, 10, 22)
		{
			v, _ := heap.Peek()
			if v != 22 {
				t.Fatalf("Expected 22, got %d", v)
			}
		}
		{
			v, _ := heap.Peek()
			if v != 22 {
				t.Fatalf("Expected 22, got %d", v)
			}
		}
	})

	t.Run("Single get", func(t *testing.T) {
		heap := setUp(1, 2)

		v, _ := heap.Get()
		if v != 2 {
			t.Fatalf("Expected 2, got %d", v)
		}
		if heap.Size() > 0 {
			t.Fatalf("expected empty")
		}
	})
	t.Run("Get removes the root", func(t *testing.T) {
		values := []int{2, 15, 3, 10, 1}
		heap := setUp(5, values...)

		{
			root, _ := heap.Get()
			if root != 15 {
				t.Fatalf("Expected 15, got %d", root)
			}
			if siz := len(values) - 1; heap.Size() != siz {
				t.Fatalf("Expected size %d, got %d", siz, heap.Size())
			}
		}
		{
			root, _ := heap.Get()
			if root != 10 {
				t.Fatalf("Expected %d, got %d", 10, root)
			}
		}
	})
}
