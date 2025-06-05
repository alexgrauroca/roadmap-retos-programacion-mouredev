package queue_test

import (
	"github.com/stretchr/testify/assert"
	"roadmap-retos-programacion-mouredev/queue"
	"testing"
)

func Test_QueueEnqueue(t *testing.T) {
	tests := []struct {
		name        string
		items       []string
		q           queue.Queue[string]
		expectedLen int
	}{
		{
			name:        "when any items enqueued, then the size is 0",
			items:       nil,
			q:           queue.New[string](0),
			expectedLen: 0,
		},
		{
			name:        "when one item enqueued, then the size is 1",
			items:       []string{"item1"},
			q:           queue.New[string](1),
			expectedLen: 1,
		},
		{
			name:        "when multiple items enqueued, then the size is equal to the number of items",
			items:       []string{"item1", "item2", "item3"},
			q:           queue.New[string](3),
			expectedLen: 3,
		},
		{
			name:        "when adding an extra item in a full queue, the last item is not added and the size remains the same",
			items:       []string{"item1", "item2", "item3", "item4", "item5"},
			q:           queue.New[string](4),
			expectedLen: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, item := range tt.items {
				tt.q.Enqueue(item)
			}
			assert.Equal(t, tt.expectedLen, tt.q.Size())
		})
	}
}

func Test_QueueDequeue(t *testing.T) {
	tests := []struct {
		name string
		q    queue.Queue[string]
		test func(t *testing.T, q queue.Queue[string])
	}{
		{
			name: "when the queue is empty, then any items are dequeued",
			q:    queue.New[string](0),
			test: func(t *testing.T, q queue.Queue[string]) {
				_, ok := q.Dequeue()

				assert.False(t, ok)
				assert.Equal(t, 0, q.Size())
			},
		},
		{
			name: "when an item is dequeued, then the item is returned and the queue has 1 less item",
			q:    queue.New[string](2),
			test: func(t *testing.T, q queue.Queue[string]) {
				q.Enqueue("item1")
				q.Enqueue("item2")
				item, ok := q.Dequeue()

				assert.True(t, ok)
				assert.Equal(t, "item1", item)
				assert.Equal(t, 1, q.Size())
			},
		},
		{
			name: "when the queue is full and item is dequeued, then a new item can be enqueued",
			q:    queue.New[string](2),
			test: func(t *testing.T, q queue.Queue[string]) {
				q.Enqueue("item1")
				q.Enqueue("item2")
				assert.Equal(t, 2, q.Size())

				_, ok := q.Dequeue()
				assert.True(t, ok)
				assert.Equal(t, 1, q.Size())

				ok = q.Enqueue("item3")
				assert.True(t, ok)
				assert.Equal(t, 2, q.Size())
			},
		},
		{
			name: "when all items are dequeued, then the queue is empty",
			q:    queue.New[string](2),
			test: func(t *testing.T, q queue.Queue[string]) {
				q.Enqueue("item1")
				q.Enqueue("item2")
				assert.Equal(t, 2, q.Size())

				item, _ := q.Dequeue()
				assert.Equal(t, "item1", item)

				item, _ = q.Dequeue()
				assert.Equal(t, "item2", item)
				assert.Equal(t, 0, q.Size())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t, tt.q)
		})
	}
}

func Test_QueuePeek(t *testing.T) {
	tests := []struct {
		name string
		q    queue.Queue[string]
		test func(t *testing.T, q queue.Queue[string])
	}{
		{
			name: "when the queue is empty, then any items are peeked",
			q:    queue.New[string](0),
			test: func(t *testing.T, q queue.Queue[string]) {
				_, ok := q.Peek()

				assert.False(t, ok)
				assert.Equal(t, 0, q.Size())
			},
		},
		{
			name: "when an item is peeked, then the item is returned and the queue has the same items",
			q:    queue.New[string](2),
			test: func(t *testing.T, q queue.Queue[string]) {
				q.Enqueue("item1")
				q.Enqueue("item2")
				item, ok := q.Peek()

				assert.True(t, ok)
				assert.Equal(t, "item1", item)
				assert.Equal(t, 2, q.Size())
			},
		},
		{
			name: "when the queue is full and item is peeked, then a new item cannot be enqueued",
			q:    queue.New[string](2),
			test: func(t *testing.T, q queue.Queue[string]) {
				q.Enqueue("item1")
				q.Enqueue("item2")
				assert.Equal(t, 2, q.Size())

				_, ok := q.Peek()
				assert.True(t, ok)
				assert.Equal(t, 2, q.Size())

				ok = q.Enqueue("item3")
				assert.False(t, ok)
			},
		},
		{
			name: "when peek is done multiple times against the same queue, then same item is returned all the times",
			q:    queue.New[string](2),
			test: func(t *testing.T, q queue.Queue[string]) {
				q.Enqueue("item1")
				q.Enqueue("item2")
				assert.Equal(t, 2, q.Size())

				item, _ := q.Peek()
				assert.Equal(t, "item1", item)

				item, _ = q.Peek()
				assert.Equal(t, "item1", item)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t, tt.q)
		})
	}
}
