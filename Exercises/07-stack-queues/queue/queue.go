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

type queue[T any] struct {
	items    []T
	capacity int
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

	q.items[q.tail] = item
	return true
}

func (q *queue[T]) Dequeue() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (q *queue[T]) Peek() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (q *queue[T]) Size() int {
	return len(q.items)
}

func (q *queue[T]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (q *queue[T]) IsFull() bool {
	return q.Size() == q.capacity
}
