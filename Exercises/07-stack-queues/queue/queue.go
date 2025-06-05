package queue

type Queue[T any] interface {
	// Enqueue an item into the queue. Returns if the operation was successful.
	Enqueue(item T) bool
	// Dequeue the first item and remove it from the queue. Returns the item and if the operation was successful.
	Dequeue() (T, bool)
	// Peek the first item from the queue, without removing it. Returns the item and if the operation was successful.
	Peek() (T, bool)
	// Size of the queue
	Size() int
	// IsEmpty validates if the queue is empty
	IsEmpty() bool
	// IsFull validates if the queue is full
	IsFull() bool
}

// queue is a circular queue, with a better memory usage
type queue[T any] struct {
	items    []T
	capacity int
	size     int
	head     int
	tail     int
}

func New[T any](capacity int) Queue[T] {
	return &queue[T]{
		items:    make([]T, capacity),
		capacity: capacity,
	}
}

func (q *queue[T]) Enqueue(item T) bool {
	if q.IsFull() {
		return false
	}

	// Insert the item to the tail of the queue
	q.items[q.tail] = item
	// Increasing the queue size by 1
	q.size++
	// Circular queue, if tail was the latest item of the slice, the next value will be 0
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	// Get the item at the head of the queue
	item := q.items[q.head]
	// Reduce the queue size by 1
	q.size--
	// Circular queue, if head was the latest item of the slice, the next value will be 0
	q.head = (q.head + 1) % q.capacity
	return item, true
}

func (q *queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	// Get the item at the head of the queue
	return q.items[q.head], true
}

func (q *queue[T]) Size() int {
	return q.size
}

func (q *queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *queue[T]) IsFull() bool {
	return q.Size() == q.capacity
}
