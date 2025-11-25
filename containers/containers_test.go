package containers

import "testing"

func TestHeap(t *testing.T) {
	mustSetUp := func(maxSize int, values ...int) *Heap {
		h, err := NewHeap(maxSize, values...)
		if err != nil {
			t.Fatalf("Failed to create heap %s", err)
		}

		return h
	}

	t.Run("Can add elements up to the max size", func(t *testing.T) {
		var (
			capacity = 1
			h        = mustSetUp(capacity)
		)

		err := h.Add(10)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if err := h.Add(15); err == nil {
			t.Fatal("Expected an error")
		}

		if siz := h.Size(); siz != capacity {
			t.Fatalf("expected an full stack (with %d elements), but it has %d elements", capacity, siz)
		}
	})

	t.Run("Peek returns but doesn't remove", func(t *testing.T) {
		heap := mustSetUp(5, 3, 10, 22)
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

		if siz := heap.Size(); siz != 3 {
			t.Fatalf("Expected 3 elements, got %d", siz)
		}
	})

	t.Run("Get removes the root", func(t *testing.T) {
		values := []int{2, 15, 3, 10, 1}
		heap := mustSetUp(5, values...)

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
