package containers

import "errors"

type queuenode[T comparable] struct {
	val        T
	prev, next *queuenode[T]
}

func mkChain[T comparable](head T, tail []T) (*queuenode[T], *queuenode[T]) {
	h := &queuenode[T]{val: head}
	cur := h
	for _, v := range tail {
		nod := &queuenode[T]{val: v, prev: cur}
		cur.next = nod
		cur = nod
	}
	return h, cur
}

var ErrEmptyQueue = errors.New("empty queue")

// Queue: A FIFO structure
type Queue[T comparable] struct {
	head, last *queuenode[T]
}

func NewQueue[T comparable](vs ...T) Queue[T] {
	if len(vs) == 0 {
		return Queue[T]{}
	}

	first, last := mkChain(vs[0], vs[1:])
	return Queue[T]{
		head: first,
		last: last,
	}

}

func (q *Queue[T]) Push(e T) {
	if q.head == nil {
		nod := &queuenode[T]{val: e}
		q.head = nod
		q.last = nod
		return
	}

	nod := &queuenode[T]{val: e, prev: q.last}
	q.last.next = nod
	q.last = nod
}

func (q *Queue[T]) Peek() (T, error) {
	var e T
	if q.head == nil || q.last == nil {
		return e, ErrEmptyQueue
	}

	return q.head.val, nil
}

func (q *Queue[T]) PeekLast() (T, error) {
	var e T
	if q.head == nil || q.last == nil {
		return e, ErrEmptyQueue
	}

	return q.last.val, nil
}

func (q *Queue[T]) Pop() (T, error) {
	if q.head == nil {
		var empty T
		return empty, ErrEmptyQueue
	}

	v := q.head.val
	q.head = q.head.next

	return v, nil
}

func (q *Queue[T]) Size() int {
	siz := 0
	for c := q.head; c != nil; c = c.next {
		siz++
	}

	return siz
}
