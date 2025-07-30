package stack

import "testing"

func TestStack_IsEmpty(t *testing.T) {
	s := Stack[int]{}
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it is not")
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("Expected stack to not be empty, but it is")
	}
}

func TestStack(t *testing.T) {
	s := Stack[int]{}
	if s.Size() != 0 {
		t.Errorf("Expected stack size to be 0, but got %d", s.Size())
	}

	s.Push(1)
	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1, but got %d", s.Size())
	}

	if s.Peek() != 1 {
		t.Errorf("Expected peeked item to be 1, but got %d", s.Peek())
	}

	item := s.Pop()
	if item != 1 {
		t.Errorf("Expected popped item to be 1, but got %d", item)
	}

	if s.Size() != 0 {
		t.Errorf("Expected stack size to be 0 after popping the only item, but got %d", s.Size())
	}

	s.Push(2)

	s.Clear()
	if s.Size() != 0 {
		t.Errorf("Expected stack size to be 0 after clearing, but got %d", s.Size())
	}

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping the only item, but it is not")
	}
}
