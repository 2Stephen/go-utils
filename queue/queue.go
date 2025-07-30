package queue

import _interface "utils/interface"

type IQueue[T any] interface {
	_interface.IBase[T]
}

type Queue[T any] struct {
	items []T
	size  uint
}

func (q *Queue[T]) Peek() T {
	if q == nil || q.size == 0 {
		var zeroValue T
		return zeroValue // Return zero value of type T if queue is empty
	}
	return q.items[0] // Return the first item in the queue
}

func (q *Queue[T]) Pop() T {
	if q == nil || q.size == 0 {
		var zeroValue T
		return zeroValue // Return zero value of type T if queue is empty
	}
	item := q.items[0]    // Get the first item
	q.items = q.items[1:] // Remove the first item from the queue
	q.size--              // Decrease the size of the queue
	return item           // Return the popped item
}

func (q *Queue[T]) Push(item T) bool {
	if q == nil {
		return false
	}
	q.items = append(q.items, item) // Add the item to the end of the queue
	q.size++                        // Increase the size of the queue
	return true                     // Return true indicating success
}

func (q *Queue[T]) Size() uint {
	if q == nil {
		return 0 // Return 0 if the queue is nil
	}
	return q.size // Return the current size of the queue
}

func (q *Queue[T]) IsEmpty() bool {
	if q == nil {
		return true // Consider a nil queue as empty
	}
	return q.size == 0 // Return true if size is 0, indicating the queue is empty
}

func (q *Queue[T]) Clear() bool {
	if q == nil {
		return false // Cannot clear a nil queue
	}
	q.items = []T{} // Reset the items slice to an empty slice
	q.size = 0      // Reset the size to 0
	return true     // Return true indicating success
}

func (q *Queue[T]) ToSlice() []T {
	if q == nil {
		return nil // Return nil if the queue is nil
	}
	return q.items // Return the items in the queue as a slice
}
