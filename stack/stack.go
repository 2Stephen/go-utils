package stack

import _interface "utils/interface"

type IStack[T any] interface {
	_interface.IBase[T]
}

type Stack[T any] struct {
	items []T
	size  uint
}

func (s Stack[T]) Peek() T {
	if s.size == 0 {
		var zeroValue T
		return zeroValue // Return zero value of type T if stack is empty
	}
	return s.items[s.size-1] // Return the last item in the stack
}

func (s Stack[T]) Pop() T {
	if s.size == 0 {
		var zeroValue T
		return zeroValue // Return zero value of type T if stack is empty
	}
	item := s.items[s.size-1]    // Get the last item
	s.items = s.items[:s.size-1] // Remove the last item from the stack
	s.size--                     // Decrease the size of the stack
	return item                  // Return the popped item
}

func (s Stack[T]) Push(item T) bool {
	if s.size == 0 {
		s.items = []T{item} // Initialize the stack with the first item
	} else {
		s.items = append(s.items, item) // Add the item to the end of the stack
	}
	s.size++    // Increase the size of the stack
	return true // Return true indicating success
}

func (s Stack[T]) Size() uint {
	if s.size == 0 {
		return 0 // Return 0 if the stack is empty
	}
	return s.size // Return the current size of the stack
}

func (s Stack[T]) IsEmpty() bool {
	if s.size == 0 {
		return true // Return true if size is 0, indicating the stack is empty
	}
	return false // Return false if the stack has items
}

func (s Stack[T]) Clear() bool {
	if s.size == 0 {
		return false // Cannot clear an empty stack
	}
	s.items = []T{} // Reset the items slice to an empty slice
	s.size = 0      // Reset the size to 0
	return true     // Return true indicating success
}

func (s Stack[T]) ToSlice() []T {
	if s.size == 0 {
		return []T{} // Return an empty slice if the stack is empty
	}
	return s.items // Return the items in the stack as a slice
}
