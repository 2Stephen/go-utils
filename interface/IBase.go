package _interface

type IBase[T any] interface {
	Peek() T
	Pop() T
	Push(item T) bool
	Size() uint
	IsEmpty() bool
	Clear() bool
	ToSlice() []T
}
