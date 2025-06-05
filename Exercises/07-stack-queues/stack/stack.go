package stack

type Stack[T any] interface {
	Push(T) bool
	Pop() (T, bool)
	Peek() (T, bool)
	IsEmpty() bool
	IsFull() bool
	Size() int
}

type stack[T any] struct {
	items    []T
	capacity int
	size     int
	top      int
}

func New[T any](capacity int) Stack[T] {
	return &stack[T]{
		items:    make([]T, capacity),
		capacity: capacity,
	}
}

func (s *stack[T]) Push(t T) bool {
	if s.IsFull() {
		return false
	}

	// Circular stack, set the next top position. If it was the latest slice position, then it will be 0
	s.top = (s.top + 1) % s.capacity
	// Add the item into the top
	s.items[s.top] = t
	// Increase stack size by 1
	s.size++
	return true
}

func (s *stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	// Fetch the item
	item := s.items[s.top]
	// Reduce the top position by 1. If it was at the beginning of the slice, then it will move to the end
	s.top = (s.top - 1 + s.capacity) % s.capacity
	// Reduce the stack size by 1
	s.size--
	return item, true
}

func (s *stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.items[s.top], true
}

func (s *stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stack[T]) IsFull() bool {
	return s.Size() == s.capacity
}

func (s *stack[T]) Size() int {
	return s.size
}
