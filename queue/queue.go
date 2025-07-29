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
	q.items = append(q.items, item)
	q.size++
	return true
}

func (q *Queue[T]) Size() uint {
	if q == nil {
		return 0
	}
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	if q == nil {
		return true
	}
	return q.size == 0
}

func (q *Queue[T]) Clear() bool {
	if q == nil {
		return false
	}
	q.items = nil
	q.size = 0
	return true
}

func (q *Queue[T]) Contains(item T) bool {
	return false
}

func (q *Queue[T]) ToSlice() []T {
	if q == nil {
		return nil
	}
	slice := make([]T, q.size)
	copy(slice, q.items)
	return slice
}
