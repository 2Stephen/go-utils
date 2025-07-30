package queue

import "testing"

func TestQueue_IsEmpty(t *testing.T) {
	q := Queue[int]{}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it is not")
	}
	q.Push(1)
	if q.IsEmpty() {
		t.Errorf("Expected queue to not be empty, but it is")
	}
}

func TestQueue(t *testing.T) {
	q := Queue[int]{}
	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0, but got %d", q.Size())
	}

	q.Push(1)
	if q.Size() != 1 {
		t.Errorf("Expected queue size to be 1, but got %d", q.Size())
	}

	if q.Peek() != 1 {
		t.Errorf("Expected peeked item to be 1, but got %d", q.Peek())
	}

	item := q.Pop()
	if item != 1 {
		t.Errorf("Expected popped item to be 1, but got %d", item)
	}

	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0 after popping the only item, but got %d", q.Size())
	}

	q.Push(2)

	q.Clear()
	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0 after clearing, but got %d", q.Size())
	}

	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty after popping the only item, but it is not")
	}
}
