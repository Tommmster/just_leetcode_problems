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

func TestQueue(t *testing.T) {

	t.Parallel()
	t.Run("Peeking an empty queue is an error", func(t *testing.T) {
		q := NewQueue[int32]()
		if _, err := q.Peek(); err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("Peek returns the first inserted value", func(t *testing.T) {
		q := NewQueue[int32]()
		q.Push(3)
		if v, _ := q.Peek(); v != 3 {
			t.Fatalf("Unexpected value: %d", v)
		}

		q.Push(5)
		if v, _ := q.Peek(); v != 3 {
			t.Fatalf("Unexpected value: %d", v)
		}
	})

	t.Run("Peek Last returns the most recent inserted value", func(t *testing.T) {
		q := NewQueue[int32](11, 12)
		if last, _ := q.PeekLast(); last != 12 {
			t.Fatalf("Unexpected value: %d", last)
		}

		if siz := q.Size(); siz != 2 {
			t.Fatalf("Unexpected size: %d", siz)
		}
	})

	t.Run("Pop removes the first value", func(t *testing.T) {
		q := NewQueue[int32](3, 5)

		if h, _ := q.Pop(); h != 3 {
			t.Fatalf("Expected the first value, got %d", h)
		}

		if h, _ := q.Pop(); h != 5 {
			t.Fatalf("Expected the first value, got %d", h)
		}

		if _, err := q.Pop(); err == nil {
			t.Fatal("Expected an error")
		}
	})
}
