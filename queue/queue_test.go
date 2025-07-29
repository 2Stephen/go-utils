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
